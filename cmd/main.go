package main

import (
	"net/http"
	"safety/controllers"
	"safety/models"
)

func main() {
	println("Application Started")

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()
	server.ListenAndServe()
}
