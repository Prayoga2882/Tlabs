package middleware

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-postgres-menu/helper"
	"log"
	"os"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		helper.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		helper.Panic(err)
	}
	return db
}
