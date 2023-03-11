package main

import "fmt"

func main() {
	x := foo(1,2,3,4,5,6,7,8,9)
	fmt.Println(x)

	yi := []int{2,3,4,5,6,7,8,9, 2034}
	y := foo(yi...)
	fmt.Println(y)

	z := variadic("Leon",4,5,6)
	fmt.Println(z)
}

func variadic(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)   // slice of int
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// Adding all the integers
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// Function takes in an unlimited number of integers (slice of int)
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)   // slice of int

	// Adding all the integers
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("The sum is",sum)
	return sum
}