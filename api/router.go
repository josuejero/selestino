// api/router.go

package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitializeRouter initializes the API routes
func InitializeRouter() *mux.Router {
	router := mux.NewRouter()

	// Define your API routes here
	router.HandleFunc("/recipes", GetRecipes).Methods("GET")

	return router
}

// GetRecipes handles the GET /recipes endpoint
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are some recipes"))
}
