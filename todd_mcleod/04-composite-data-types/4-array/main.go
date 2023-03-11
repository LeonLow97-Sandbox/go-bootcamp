package main

import "fmt"

func main() {
	// Creating an array of size 5
	// Array are of a single type
	// length of array is part of the Array's type
	var x [5]int
	fmt.Println(x)

	// Assigning a value of '42' at position 3 of the array
	x[3] = 42
	fmt.Println(x)

	// Getting the length of array
	fmt.Println(len(x))
}