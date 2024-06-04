// cmd/main.go

package main

import (
	"log"
	"net/http"

	"github.com/josuejero/selestino/api"
	"github.com/josuejero/selestino/pkg/config"
)

func main() {
	db := config.Connect()
	defer db.Close()

	router := api.InitializeRouter(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
