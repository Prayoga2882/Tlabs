package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

var (
	secretkey = "secretkeyjwt"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomInt() int {
	return rand.Int()
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(username string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		_ = fmt.Errorf("something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
