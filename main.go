package main

import (
	"fmt"
	"go-postgres-menu/controllers"
	"go-postgres-menu/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	// API Routes
	r.HandleFunc("/api/menu", controllers.CreateMenu).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/menu", controllers.GetAllMenu).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/menu/{id}", controllers.GetMenu).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/menu/{id}", controllers.UpdateMenu).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/menu/{id}", controllers.DeleteMenu).Methods("DELETE", "OPTIONS")

	fmt.Println("Starting server on 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
