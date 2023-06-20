package router

import (
	"github.com/gorilla/mux"
	"main/controllers"
	"main/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/signin", controllers.SignIn).Methods("POST")
	// API Routes
	router.HandleFunc("/api/menu", middleware.IsAuthorized(controllers.CreateMenu)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/menu", controllers.GetAllMenu).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/menu/{id}", controllers.GetMenu).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/menu/{id}", middleware.IsAuthorized(controllers.UpdateMenu)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/menu/{id}", middleware.IsAuthorized(controllers.DeleteMenu)).Methods("DELETE", "OPTIONS")
	return router
}
