package controller

import (
	"encoding/json"
	"net/http"

	"my-app-server/helpers"
	"my-app-server/services"
	"my-app-server/types"

	"github.com/gorilla/mux"
)

func getAllUsersHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "GET" {
			users, err := services.GetUsers()
			if err == nil {
				helpers.RespondWithSuccess(users, writer)
			} else {
				helpers.RespondWithError(err, writer)
			}
		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func createUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "POST" {
			var user types.User
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
		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func updateUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "PUT" {
			var user types.User
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
		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getOneUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "GET" {
			idAsString := mux.Vars(request)["id"]
			id, err := helpers.StringToInt64(idAsString)
			if err != nil {
				helpers.RespondWithError(err, writer)
				// We return, so we stop the function flow
				return
			}
			user, err := services.GetUserById(id)
			if err != nil {
				helpers.RespondWithError(err, writer)
			} else {
				helpers.RespondWithSuccess(user, writer)
			}
		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func deleteUserHandler(router *mux.Router) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "DELETE" {
			idAsString := mux.Vars(request)["id"]
			id, err := helpers.StringToInt64(idAsString)
			if err != nil {
				helpers.RespondWithError(err, writer)
				// We return, so we stop the function flow
				return
			}
			err = services.DeleteUser(id)
			if err != nil {
				helpers.RespondWithError(err, writer)
			} else {
				helpers.RespondWithSuccess(true, writer)
			}
		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func SetupRoutesForUser(router *mux.Router) {
	// Create user
	router.HandleFunc("/user", helpers.ValidateTokenMiddleware(getAllUsersHandler(router))).Methods(http.MethodGet)
	router.HandleFunc("/user", helpers.ValidateTokenMiddleware(createUserHandler(router))).Methods(http.MethodPost)
	router.HandleFunc("/user", helpers.ValidateTokenMiddleware(updateUserHandler(router))).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", helpers.ValidateTokenMiddleware(getOneUserHandler(router))).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", helpers.ValidateTokenMiddleware(deleteUserHandler(router))).Methods(http.MethodDelete)

}
