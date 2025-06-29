package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func a(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This represents all routes.")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is a dog route.")
}

func m(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", "Leon")

	// io.WriteString(res, "This is a route for me, Leon!")
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
