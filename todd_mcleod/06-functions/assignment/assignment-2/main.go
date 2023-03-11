package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7732}
	fmt.Println(foo(s...))

	fmt.Println(bar(s))
}

// Takes in a variadic parameter of type int
// Pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
func foo(v ...int) int {
	sum := 0
	for _, v := range v {
		sum += v
	}
	return sum
}

// Takes in a parameter of type []int
// returns the sum of all values of type int passed in
func bar(si []int) int {
	sum := 0
	for _, v := range si {
		sum += v
	}
	return sum
}