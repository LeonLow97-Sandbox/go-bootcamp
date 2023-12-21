package main

import (
	"log"
	"net/http"
)

type application struct {
}

func main() {
	// Create an instance of your application
	app := &application{} // Assuming you have initialized your application struct

	// Retrieve the router with routes configured
	router := app.routes()

	// Define the server address and port
	addr := ":8080" // Change this to the desired port number

	// Start the HTTP server
	log.Printf("Server is running on %s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
