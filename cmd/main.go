// cmd/main.go

package main

import (
	"net/http"

	"github.com/josuejero/selestino/api"
	"github.com/josuejero/selestino/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	db := config.Connect()
	defer db.Close()

	config.InitRedis()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	router := api.InitializeRouter(db)

	logrus.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logrus.Fatalf("could not start server: %v", err)
	}
}
