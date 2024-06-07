package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josuejero/selestino/api"
	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRecipeRepo *repository.MockRecipeRepository

func setup() {
	mockRecipeRepo = &repository.MockRecipeRepository{}
	api.SetRecipeRepo(mockRecipeRepo)
}

func TestGetRecipes(t *testing.T) {
	setup()

	mockRecipeRepo.On("GetAllRecipes").Return([]models.Recipe{
		{Name: "Ceviche", Ingredients: "Fish, Lemon, Salt", Instructions: "Mix ingredients and serve"},
	}, nil)

	req, _ := http.NewRequest("GET", "/recipes", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetRecipes)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var recipes []models.Recipe
	err := json.NewDecoder(rr.Body).Decode(&recipes)
	assert.NoError(t, err)
	assert.NotEmpty(t, recipes)
}

func TestAddRecipe(t *testing.T) {
	setup()

	mockRecipeRepo.On("AddRecipe", mock.AnythingOfType("models.Recipe")).Return(nil)

	recipe := models.Recipe{Name: "Test Recipe", Ingredients: "Test Ingredients", Instructions: "Test Instructions"}
	payload, _ := json.Marshal(recipe)
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.AddRecipe)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestSearchRecipesByCriteria(t *testing.T) {
	setup()

	mockRecipeRepo.On("SearchRecipesByCriteria", mock.AnythingOfType("map[string]string")).Return([]models.Recipe{
		{Name: "Test Recipe", Ingredients: "Test Ingredients", Instructions: "Test Instructions"},
	}, nil)

	req, _ := http.NewRequest("GET", "/recipes/search?ingredients=Test+Ingredients", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.SearchRecipesByCriteria)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var recipes []models.Recipe
	err := json.NewDecoder(rr.Body).Decode(&recipes)
	assert.NoError(t, err)
	assert.NotEmpty(t, recipes)
}
