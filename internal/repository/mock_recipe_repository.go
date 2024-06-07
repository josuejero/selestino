// internal/repository/mock_recipe_repository.go

package repository

import (
	"github.com/josuejero/selestino/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockRecipeRepository struct {
	mock.Mock
}

func (m *MockRecipeRepository) GetAllRecipes() ([]models.Recipe, error) {
	args := m.Called()
	return args.Get(0).([]models.Recipe), args.Error(1)
}

func (m *MockRecipeRepository) AddRecipe(recipe models.Recipe) error {
	args := m.Called(recipe)
	return args.Error(0)
}

func (m *MockRecipeRepository) SearchRecipesByCriteria(criteria map[string]string) ([]models.Recipe, error) {
	args := m.Called(criteria)
	return args.Get(0).([]models.Recipe), args.Error(1)
}
