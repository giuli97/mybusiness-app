package routes

import (
	"encoding/json"
	"net/http"

	controller "my-app-server/controller"
	helpers "my-app-server/helpers"
	"my-app-server/types"

	"github.com/gorilla/mux"
)

func SetupRoutesForUsers(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	helpers.EnableCORS(router)

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := controller.GetUsers()
		if err == nil {
			helpers.RespondWithSuccess(users, w)
		} else {
			helpers.RespondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := helpers.StringToInt64(idAsString)
		if err != nil {
			helpers.RespondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		user, err := controller.GetUserById(id)
		if err != nil {
			helpers.RespondWithError(err, w)
		} else {
			helpers.RespondWithSuccess(user, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var user types.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			helpers.RespondWithError(err, w)
		} else {
			err := controller.CreateUser(user)
			if err != nil {
				helpers.RespondWithError(err, w)
			} else {
				helpers.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var user types.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			helpers.RespondWithError(err, w)
		} else {
			err := controller.UpdateUser(user)
			if err != nil {
				helpers.RespondWithError(err, w)
			} else {
				helpers.RespondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := helpers.StringToInt64(idAsString)
		if err != nil {
			helpers.RespondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		err = controller.DeleteUser(id)
		if err != nil {
			helpers.RespondWithError(err, w)
		} else {
			helpers.RespondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

}
