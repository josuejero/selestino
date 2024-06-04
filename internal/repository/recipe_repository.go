// internal/repository/recipe_repository.go

package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/josuejero/selestino/internal/models"
)

type RecipeRepository struct {
	DB *sql.DB
}

func (r *RecipeRepository) GetAllRecipes() ([]models.Recipe, error) {
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
	return recipes, nil
}

func (r *RecipeRepository) AddRecipe(recipe models.Recipe) error {
	_, err := r.DB.Exec("INSERT INTO recipes (name, ingredients, instructions) VALUES ($1, $2, $3)",
		recipe.Name, recipe.Ingredients, recipe.Instructions)
	return err
}

func (r *RecipeRepository) SearchRecipesByIngredients(ingredients []string) ([]models.Recipe, error) {
	query := "SELECT id, name, ingredients, instructions FROM recipes WHERE"

	var conditions []string
	for _, ingredient := range ingredients {
		conditions = append(conditions, fmt.Sprintf(" ingredients LIKE '%%%s%%'", ingredient))
	}
	query += strings.Join(conditions, " AND ")

	rows, err := r.DB.Query(query)
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
