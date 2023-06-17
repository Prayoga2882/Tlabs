package main

import (
	"fmt"
	"go-postgres-menu/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on 0.0.0.0:8080")

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
