package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// create a file
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Leon Low")

	io.Copy(f, r)
}
