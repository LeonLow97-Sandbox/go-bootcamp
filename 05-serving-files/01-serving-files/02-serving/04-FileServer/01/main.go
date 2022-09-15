package main

import (
	"io"
	"net/http"
)

func main() {
	// func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
	// takes all the files in the current directory
	http.Handle("/", http.FileServer((http.Dir("."))))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}
