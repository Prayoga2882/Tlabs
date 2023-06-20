package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"main/helper"
	"net/http"
	"os"
)

var (
	router    *mux.Router
	secretkey = "secretkeyjwt"
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

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] == nil {
			var err helper.Error
			err = helper.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing token")
			}
			return mySigningKey, nil
		})
		if err != nil {
			var err helper.Error
			err = helper.SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			var err helper.Error
			err = helper.SetError(err, "Not Authorized.")
			json.NewEncoder(w).Encode(err)
			return
		}

		handler.ServeHTTP(w, r)
	}
}
