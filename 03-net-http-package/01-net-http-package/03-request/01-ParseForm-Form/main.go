package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// func (r *Request) ParseForm() error
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// Form contains the parsed form data and the url
	// This field is only available after the ParseForm is called.
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
