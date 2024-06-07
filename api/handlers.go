// api/handlers.go

package api

import (
	"encoding/json"
	"net/http"

	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/internal/repository"
)

var recipeRepo repository.RecipeRepositoryInterface

func SetRecipeRepo(r repository.RecipeRepositoryInterface) {
	recipeRepo = r
}

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := recipeRepo.GetAllRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func AddRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := recipeRepo.AddRecipe(recipe); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchRecipesByCriteria(w http.ResponseWriter, r *http.Request) {
	criteria := make(map[string]string)
	for key, values := range r.URL.Query() {
		criteria[key] = values[0]
	}

	recipes, err := recipeRepo.SearchRecipesByCriteria(criteria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
