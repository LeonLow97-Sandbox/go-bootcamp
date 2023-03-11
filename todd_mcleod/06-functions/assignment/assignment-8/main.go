package main

import "fmt"

func main() {
	foo()()

	b := bar()
	fmt.Println(b())
}

// Create a func which returns a func
// Assign the returned func to a variable
// Call the returned func
func foo() func() {
	x := func() {
		fmt.Println("Hello Leon")
	}
	return x
}

func bar() func() int {
	return func() int {
		return 42
	}
}