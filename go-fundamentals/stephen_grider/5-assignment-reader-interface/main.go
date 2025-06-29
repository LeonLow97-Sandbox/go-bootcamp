package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Access to the file name from command line argument list
	// fmt.Println(os.Args[1])

	// this 'f' implemeted the Reader Interface
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Copy function (writer, reader)
	io.Copy(os.Stdout, f)
}
