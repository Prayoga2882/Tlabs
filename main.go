package main

import (
	"go-postgres-menu/router"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := router.Router()
	var port = envPortOr("3000")
	log.Println("Starting server on port " + "0.0.0.0" + port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
