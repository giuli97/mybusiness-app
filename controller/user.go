package controller

import (
	"encoding/json"
	"net/http"

	"my-app-server/helpers"
	"my-app-server/models"
	"my-app-server/services"

	"github.com/gorilla/mux"
)

func getAllUsersHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		users, err := services.GetUsers()
		if err == nil {
			helpers.RespondWithSuccess(users, writer)
		} else {
			helpers.RespondWithError(err, writer)
		}
	}
}

func createUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		var user models.User
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			helpers.RespondWithError(err, writer)
		} else {
			err := services.CreateUser(user)
			if err != nil {
				helpers.RespondWithError(err, writer)
			} else {
				helpers.RespondWithSuccess(true, writer)
			}
		}
	}
}

func updateUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		var user models.User
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			helpers.RespondWithError(err, writer)
		} else {
			err := services.UpdateUser(user)
			if err != nil {
				helpers.RespondWithError(err, writer)
			} else {
				helpers.RespondWithSuccess(true, writer)
			}
		}
	}
}

func getOneUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		idAsString := mux.Vars(request)["id"]
		id, err := helpers.StringToInt64(idAsString)
		if err != nil {
			helpers.RespondWithError(err, writer)
			return
		}
		user, err := services.GetUserById(id)
		if err != nil {
			helpers.RespondWithError(err, writer)
		} else {
			helpers.RespondWithSuccess(user, writer)
		}
	}
}

func deleteUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		idAsString := mux.Vars(request)["id"]
		id, err := helpers.StringToInt64(idAsString)
		if err != nil {
			helpers.RespondWithError(err, writer)
			return
		}
		err = services.DeleteUser(id)
		if err != nil {
			helpers.RespondWithError(err, writer)
		} else {
			helpers.RespondWithSuccess(true, writer)
		}
	}
}

func SetupRoutesForUser(router *mux.Router) {
	// Get all users
	router.HandleFunc("/user", helpers.ValidateTokenMiddleware(getAllUsersHandler(router))).Methods(http.MethodGet)
	// Create user
	router.HandleFunc("/user", helpers.ValidateTokenMiddleware(createUserHandler(router))).Methods(http.MethodPost)
	// Update user
	router.HandleFunc("/user", helpers.ValidateTokenMiddleware(updateUserHandler(router))).Methods(http.MethodPut)
	// Get one user
	router.HandleFunc("/user/{id}", helpers.ValidateTokenMiddleware(getOneUserHandler(router))).Methods(http.MethodGet)
	// Delete user
	router.HandleFunc("/user/{id}", helpers.ValidateTokenMiddleware(deleteUserHandler(router))).Methods(http.MethodDelete)
}
