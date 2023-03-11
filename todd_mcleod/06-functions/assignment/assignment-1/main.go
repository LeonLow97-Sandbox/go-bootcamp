package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

// Create a function with the identifier foo that returns an int
func foo() int {
	s := 5
	return s
}

// Create a function with the identifier bar that returns an int and a string
func bar() (int, string) {
	return 16, "Jie Wei"
}