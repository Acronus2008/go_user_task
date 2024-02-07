package main

import (
	"net/http"

	"api_rest_postgresql/datasource"
	"api_rest_postgresql/routes"

	"github.com/gorilla/mux"
)

func main() {
	datasource.DatasourceConnection()

	datasource.ExecuteMigrations()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
