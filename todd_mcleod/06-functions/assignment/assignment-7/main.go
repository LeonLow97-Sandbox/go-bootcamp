package main

import "fmt"

var x int

func main() {
	// Assign a func to a variable, then call the func
	y := func() int {
		x = 5
		return x
	}

	fmt.Printf("%T\n",y)
	fmt.Println(y())
}
