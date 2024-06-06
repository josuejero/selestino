// internal/repository/recipe_repository_test.go

package repository_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/internal/repository"
	"github.com/stretchr/testify/assert"
)

var repo *repository.RecipeRepository
var mock sqlmock.Sqlmock

func setupRecipeRepo(t *testing.T) func() {
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repo = &repository.RecipeRepository{DB: db}

	return func() {
		db.Close()
	}
}

func TestGetAllRecipes(t *testing.T) {
	teardown := setupRecipeRepo(t)
	defer teardown()

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
	teardown := setupRecipeRepo(t)
	defer teardown()

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

func TestSearchRecipesByCriteria(t *testing.T) {
	teardown := setupRecipeRepo(t)
	defer teardown()

	rows := sqlmock.NewRows([]string{"id", "name", "ingredients", "instructions"}).
		AddRow(1, "Ceviche", "Fish, Lemon, Salt", "Mix ingredients and serve")

	criteria := map[string]string{"ingredients": "Fish", "name": "Ceviche"}
	query := "SELECT id, name, ingredients, instructions FROM recipes WHERE 1=1 AND ingredients LIKE \\$1 AND name LIKE \\$2"
	mock.ExpectQuery(query).WithArgs("%Fish%", "%Ceviche%").WillReturnRows(rows)

	recipes, err := repo.SearchRecipesByCriteria(criteria)
	assert.NoError(t, err)
	assert.Len(t, recipes, 1)
	assert.Equal(t, "Ceviche", recipes[0].Name)
}
