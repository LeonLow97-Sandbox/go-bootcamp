package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// func ParseGlob(pattern string) (*Template, error)
	// tpl, err := template.ParseGlob("templates/*.gohtml")
	tpl, err := template.ParseGlob("templates/*") // take all the files in the current folder
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
