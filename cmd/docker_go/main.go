package main

import (
	"docker_go/internal/database"
	"docker_go/internal/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// USER Management API endpoint

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing docker!")
}

func main() {

	// init the database
	database.InitDB()

	// API endpoint using mux
	r := mux.NewRouter()

	// the difference between http and mux handlefunc is that mux handlefunc has a path prefix
	r.HandleFunc("/", helloHandler)

	// handle endpoints with versioning by grouping them for better organization
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/users/register", handlers.RegisterUser).Methods("POST")
	v1.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	//http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", r)
	fmt.Println("Server running on port 8080")
}
