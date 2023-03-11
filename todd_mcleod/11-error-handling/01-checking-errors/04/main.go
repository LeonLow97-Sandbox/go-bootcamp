package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// create a file
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
