package controller

import (
	"encoding/json"
	"my-app-server/models"
	"my-app-server/services"
	"net/http"

	"github.com/gorilla/mux"
)

func LoginHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		body := request.Body
		defer body.Close()
		var user models.User
		err := json.NewDecoder(body).Decode(&user)
		if err != nil {
			http.Error(writer, "cannot decode json", http.StatusBadRequest)
			return
		}
		// TODO: funcion validadora del usuario que le pregunte a la BD si existe
		if user.UserName == "giuli" && user.Password == "123" {
			res := services.GetTokens()
			_ = json.NewEncoder(writer).Encode(&res)
		} else {
			http.Error(writer, "bad credentials", http.StatusUnauthorized)
			return
		}
	}
}

func SetupRoutesForAuthentication(router *mux.Router) {
	// Login
	router.HandleFunc("/login", LoginHandler()).Methods(http.MethodPost)
	// TODO: sign up, logout, refresh token
}
