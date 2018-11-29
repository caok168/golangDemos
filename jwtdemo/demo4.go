package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	mySigningKey := []byte("AllYourBase")

	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 150000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}
