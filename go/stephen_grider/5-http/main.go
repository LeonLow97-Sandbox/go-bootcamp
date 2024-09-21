package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom type for Write function
type logWriter struct {}

func main() {
	resp, err := http.Get("https://www.google.com")

	// error handling
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// takes a type of slice, and number of elements that the slice can contain
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

// this function is implementing the writer interface
// custom writer interface
func (logWriter) Write(bs []byte) (int, error) {
	// return 1, nil
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
} 
