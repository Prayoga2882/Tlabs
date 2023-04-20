package main

import (
	"fmt"
	"go-postgres-menu/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
