package services

import (
	"encoding/json"
	"log"
	"my-app-server/types"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func UserHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "POST" {
			body := request.Body
			defer body.Close()
			var user types.User
			err := json.NewDecoder(body).Decode(&user)

			if err != nil {
				http.Error(writer, "cannot decode json", http.StatusBadRequest)
				return
			}

			if user.UserName == "giuli" && user.Password == "123" {
				res := getTokens()
				_ = json.NewEncoder(writer).Encode(&res)
			} else {
				http.Error(writer, "bad credentials", http.StatusUnauthorized)
				return
			}

		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getTokens() *types.JWTResponse {
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
