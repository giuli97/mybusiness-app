package controller

import (
	"my-app-server/helpers"

	"github.com/gorilla/mux"
)

func SetupAllRoutes(router *mux.Router) {
	// First enable cors
	helpers.EnableCORS(router)

	SetupRoutesForAuthentication(router)
	SetupRoutesForUser(router)
}
