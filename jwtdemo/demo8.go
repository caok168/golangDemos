package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwMCwiaXNzIjoidGVzdCJ9.bP8ZXh1bFHBci8cGx9i1ryNdk23-Am22yFwgodysKJs"

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims1{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims1); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
}

type MyCustomClaims1 struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}
