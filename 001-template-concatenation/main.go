package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Leon"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello Jie Wei!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)

	// Method 2
	// func Create(name string) (*File, error)
	nf, err := os.Create("index2.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}

	// Close file at the end
	defer nf.Close()

	// func Copy(dst Writer, src Reader) (written int64, err error)
	// Copy copies from src to dst until either EOF is reached on src or an error occurs.
	// NewReader returns a new Reader reading from s.
	_, err = io.Copy(nf, strings.NewReader(tpl))
	if err != nil {
		fmt.Println("There was an error copying")
	}
}

func thirdMethod() {
	name := os.Args[1] // name
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset = "UTF-8">
		<title>Hello Jie Wei!</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	`)

	nf, err := os.Create("index3.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
