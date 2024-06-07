package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/pkg/config"
	"github.com/olivere/elastic/v7"
)

type RecipeRepositoryInterface interface {
	GetAllRecipes() ([]models.Recipe, error)
	AddRecipe(recipe models.Recipe) error
	SearchRecipesByCriteria(criteria map[string]string) ([]models.Recipe, error)
}

type RecipeRepository struct {
	DB *sql.DB
}

var _ RecipeRepositoryInterface = &RecipeRepository{}

func (r *RecipeRepository) GetAllRecipes() ([]models.Recipe, error) {
	ctx := context.Background()
	cacheKey := "all_recipes"

	// Try to get the result from Redis cache
	cachedResult, err := config.GetRedis(ctx, cacheKey)
	if err == nil && cachedResult != "" {
		var recipes []models.Recipe
		if err := json.Unmarshal([]byte(cachedResult), &recipes); err == nil {
			return recipes, nil
		}
	}

	// If cache miss, query the database
	rows, err := r.DB.Query("SELECT id, name, ingredients, instructions FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Ingredients, &recipe.Instructions); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	// Cache the result in Redis
	serializedRecipes, err := json.Marshal(recipes)
	if err == nil {
		config.SetRedis(ctx, cacheKey, serializedRecipes)
	}

	return recipes, nil
}

func (r *RecipeRepository) AddRecipe(recipe models.Recipe) error {
	result, err := r.DB.Exec("INSERT INTO recipes (name, ingredients, instructions) VALUES ($1, $2, $3)",
		recipe.Name, recipe.Ingredients, recipe.Instructions)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Index document in Elasticsearch
	err = config.IndexDocument("recipes", strconv.FormatInt(id, 10), recipe)
	if err != nil {
		return err
	}

	// Invalidate cache
	ctx := context.Background()
	config.DelRedis(ctx, "all_recipes")

	return nil
}

func (r *RecipeRepository) SearchRecipesByCriteria(criteria map[string]string) ([]models.Recipe, error) {
	query := elastic.NewBoolQuery()
	for key, value := range criteria {
		query = query.Must(elastic.NewMatchQuery(key, value))
	}

	searchResult, err := config.ESClient.Search().
		Index("recipes").
		Query(query).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	var recipes []models.Recipe
	for _, hit := range searchResult.Hits.Hits {
		var recipe models.Recipe
		err := json.Unmarshal(hit.Source, &recipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
