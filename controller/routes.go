package controller

import (
	"my-app-server/middleware"

	"github.com/gorilla/mux"
)

func SetupAllRoutes(router *mux.Router) {
	// First enable cors
	middleware.EnableCORS(router)

	SetupRoutesForAuthentication(router)
	SetupRoutesForUser(router)
}
