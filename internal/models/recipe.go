// internal/models/recipe.go

package models

type Recipe struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
}
