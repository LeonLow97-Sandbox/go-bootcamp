package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(w, "Hello world")
	case "/dog":
		io.WriteString(w, "doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty")
	default:
		io.WriteString(w, "Invalid URL")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
