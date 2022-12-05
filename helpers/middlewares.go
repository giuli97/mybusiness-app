package helpers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ValidateTokenMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		fmt.Println("El token: ", bearerToken)
		if validateToken(bearerToken) {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
}

func validateToken(s string) bool {
	if strings.Contains(strings.ToLower(s), "bearer ") {
		// token := strings.Replace(s, "bearer ", "", 1) // handle bearer ignoring casing

		token := s[7:]
		mySigningKey := []byte("AllYourBase") // centralize this secret key
		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return mySigningKey, nil
		})

		if err != nil {
			log.Println(err.Error())
			return false
		}

		return t.Valid
	}
	return false
}
