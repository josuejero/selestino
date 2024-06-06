// internal/repository/recipe_repository.go

package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/pkg/config"
)

type RecipeRepository struct {
	DB *sql.DB
}

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
	_, err := r.DB.Exec("INSERT INTO recipes (name, ingredients, instructions) VALUES ($1, $2, $3)",
		recipe.Name, recipe.Ingredients, recipe.Instructions)

	// Invalidate cache
	ctx := context.Background()
	config.DelRedis(ctx, "all_recipes")

	return err
}

func (r *RecipeRepository) SearchRecipesByCriteria(criteria map[string]string) ([]models.Recipe, error) {
	query := "SELECT id, name, ingredients, instructions FROM recipes WHERE 1=1"

	var args []interface{}
	i := 1
	for key, value := range criteria {
		query += fmt.Sprintf(" AND %s LIKE $%d", key, i)
		args = append(args, "%"+value+"%")
		i++
	}

	rows, err := r.DB.Query(query, args...)
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

	return recipes, nil
}
