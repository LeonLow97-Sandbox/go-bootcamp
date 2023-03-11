package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	y := func(x int) {
		fmt.Println("The year I was born in is", x)
	}

	y(2003)
}