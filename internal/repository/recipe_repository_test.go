package repository_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/josuejero/selestino/internal/repository"
	"github.com/josuejero/selestino/pkg/config"
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

	config.InitRedis() // Initialize Redis client for the test

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
