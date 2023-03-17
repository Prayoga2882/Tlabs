package helper

import "math/rand"

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomInt() int {
	return rand.Int()
}
