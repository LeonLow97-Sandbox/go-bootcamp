package main

import (
	"log"
	"net/http"
)

func main() {
	// if there is an error, then it logs the error
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
