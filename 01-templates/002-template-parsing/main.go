package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Parse Files
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// func (t *Template) Execute(wr io.Writer, data any) error
	// os.Stdout implements the writer interface
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	writeToHtml()
}

func writeToHtml() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	// write to file
	// because File implements the writer interface
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
