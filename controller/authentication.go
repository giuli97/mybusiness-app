package controller

import (
	"my-app-server/services"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForAuthentication(router *mux.Router) {
	// Login
	router.HandleFunc("/login", services.UserHandler()).Methods(http.MethodPost)
	// TODO: sign up, logout, refresh token

}
