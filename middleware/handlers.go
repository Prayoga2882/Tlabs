package middleware

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
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

func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		helper.Panic(err)
	}
}

func CreateConnectionPgx() (*pgx.Conn, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		helper.Panic(err)
	}
	return conn, err
}

func CloseConnectionPgx(conn *pgx.Conn) {
	err := conn.Close(context.Background())
	if err != nil {
		helper.Panic(err)
	}
}
