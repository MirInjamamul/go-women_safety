package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{email}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	// router.HandleFunc("/api/complains", GetAllComplains).Methods("GET")
	// router.HandleFunc("/api/complain/{id}", GetComplain).Methods("GET")
	// router.HandleFunc("/api/complain", CreateComplain).Methods("POST")
	// router.HandleFunc("/api/complain/{id}", DeleteComplain).Methods("DELETE")

	return router
}
