// Entrypoint for API
package main

import (
	"log"
	"net/http"
	"os"
	"sgx/server/db"
	"sgx/server/store"

	"github.com/gorilla/handlers"
)

func main() {
	// Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Attempt the database connection before starting the server
	db.Connect()

	router := store.NewRouter() // create routes

	go store.HandleSocketMessages()

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "HEAD", "DELETE", "PUT", "OPTIONS"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}
