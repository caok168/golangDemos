package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	//token := GenToken()
	//fmt.Println(token)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDM0NTc4NjQsImlzcyI6IkJpdGNoIiwibmJmIjoxNTQzNDU2ODY0fQ.GQvuoYn886xI5UtX_Qjwl1S2iPRnxAAGOqN1u57u340"

	isok := CheckToken(token)
	fmt.Println("isok:", isok)
}

var (
	key []byte = []byte("Hello World!This is secret!")
)

// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 10),
		Issuer:    "Bitch",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return ss
}

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parse with claims failed.", err)
		return false
	}
	return true
}
