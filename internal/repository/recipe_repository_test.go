// internal/repository/recipe_repository_test.go

package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/josuejero/selestino/internal/models"
	"github.com/stretchr/testify/assert"
)

var repo *RecipeRepository
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo = &RecipeRepository{DB: db}

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func TestGetAllRecipes(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "ingredients", "instructions"}).
		AddRow(1, "Ceviche", "Fish, Lemon, Salt", "Mix ingredients and serve")

	mock.ExpectQuery("SELECT id, name, ingredients, instructions FROM recipes").
		WillReturnRows(rows)

	recipes, err := repo.GetAllRecipes()
	assert.NoError(t, err)
	assert.Len(t, recipes, 1)
	assert.Equal(t, "Ceviche", recipes[0].Name)
}

func TestAddRecipe(t *testing.T) {
	recipe := models.Recipe{
		Name:         "Ceviche",
		Ingredients:  "Fish, Lemon, Salt",
		Instructions: "Mix ingredients and serve",
	}

	mock.ExpectExec("INSERT INTO recipes").WithArgs(recipe.Name, recipe.Ingredients, recipe.Instructions).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.AddRecipe(recipe)
	assert.NoError(t, err)
}

func TestSearchRecipesByIngredients(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "ingredients", "instructions"}).
		AddRow(1, "Ceviche", "Fish, Lemon, Salt", "Mix ingredients and serve")

	ingredients := []string{"Fish", "Lemon"}
	query := "SELECT id, name, ingredients, instructions FROM recipes WHERE ingredients LIKE '%Fish%' AND ingredients LIKE '%Lemon%'"
	mock.ExpectQuery(query).WillReturnRows(rows)

	recipes, err := repo.SearchRecipesByIngredients(ingredients)
	assert.NoError(t, err)
	assert.Len(t, recipes, 1)
	assert.Equal(t, "Ceviche", recipes[0].Name)
}
