package main

import (
	"log"
	"net/http"
)

func main() {
	// if you have a file "index.html", it will response with that
	// instead of showing the different folders in the current directory
	// special case
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
