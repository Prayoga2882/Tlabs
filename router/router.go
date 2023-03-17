package router

import (
	"github.com/gorilla/mux"
	"go-postgres-stocks/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// API Routes
	router.HandleFunc("/api/stocks", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stocks/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
