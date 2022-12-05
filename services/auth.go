package services

import (
	"log"
	"my-app-server/types"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetTokens() *types.JWTResponse {
	token := signToken()
	refresh := signRefreshToken()
	return &types.JWTResponse{
		Token:   token,
		Refresh: refresh,
	}
}

func signToken() string {
	// Expiration time = current time + 12 hours
	expirationTime := time.Now().Local().Add(time.Hour * time.Duration(12))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": expirationTime,
	})
	mySigningKey := []byte("AllYourBase")
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return t
}

func signRefreshToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "baz",
		"exp":     time.Date(2020, 11, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"refresh": true,
	})
	mySigningKey := []byte("AllYourBase")
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return t
}
