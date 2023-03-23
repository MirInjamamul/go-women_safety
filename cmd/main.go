package main

import (
	"net/http"

	"go-safety/controllers"
	"go-safety/models"
)

func main() {

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
