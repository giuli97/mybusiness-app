package main

import (
	"log"
	"net/http"
	"time"

	"my-app-server/helpers"
	routes "my-app-server/routes"

	"github.com/gorilla/mux"
)

func main() {
	//err := godotenv.Load(".env")

	bd, err := helpers.GetDB()
	if err != nil {
		log.Printf("Error with database" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error making connection to DB. Please check credentials. The error is: " + err.Error())
			return
		}
	}

	// Define routes
	router := mux.NewRouter()
	routes.SetupRoutesForUsers(router)
	// .. here you can define more routes
	// ...
	// for example setupRoutesForGenres(router)

	// Setup and start server
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())

}
