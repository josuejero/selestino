// api/router.go

package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/internal/repository"
)

var recipeRepo *repository.RecipeRepository

func InitializeRouter(db *sql.DB) *mux.Router {
	recipeRepo = &repository.RecipeRepository{DB: db}
	router := mux.NewRouter()

	// Define your API routes here
	router.HandleFunc("/recipes", GetRecipes).Methods("GET")
	router.HandleFunc("/recipes", AddRecipe).Methods("POST")
	router.HandleFunc("/recipes/search", SearchRecipesByIngredients).Methods("GET")

	return router
}

// GetRecipes handles the GET /recipes endpoint
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := recipeRepo.GetAllRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

// AddRecipe handles the POST /recipes endpoint
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

// SearchRecipesByIngredients handles the GET /recipes/search endpoint
func SearchRecipesByIngredients(w http.ResponseWriter, r *http.Request) {
	ingredientsParam := r.URL.Query().Get("ingredients")
	if ingredientsParam == "" {
		http.Error(w, "ingredients query parameter is required", http.StatusBadRequest)
		return
	}

	ingredients := strings.Split(ingredientsParam, ",")

	recipes, err := recipeRepo.SearchRecipesByIngredients(ingredients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
