package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	x := foo(sum, xi...)

	fmt.Println(x)
}

func foo(f func(x ...int) int, s ...int) int {

	// To append values of `s`
	var yi []int
	for _, v := range s {
		yi = append(yi, v)
	}

	return f(yi...)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}