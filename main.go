package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jintoples/go-jwt-mux/controllers/authcontroller"
	"github.com/jintoples/go-jwt-mux/controllers/productcontroller"
	"github.com/jintoples/go-jwt-mux/middleware"
	"github.com/jintoples/go-jwt-mux/models"
)

func main() {

	models.ConnectDatabase()
	route := mux.NewRouter()

	route.HandleFunc("/login", authcontroller.Login).Methods("POST")
	route.HandleFunc("/register", authcontroller.Register).Methods("POST")
	route.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := route.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middleware.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":3000", route))
}
