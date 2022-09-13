package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.txt")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse multiple files together in a container
	// this tpl is a pointer to template
	// attach 2 more files to pointer to the template
	tpl, err = tpl.ParseFiles("two.txt", "hello.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "hello.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
