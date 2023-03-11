package main

import "fmt"

func main() {
	fmt.Println("Total Sum", mySum(2, 3, 4, 5, 6))
	fmt.Println("Total Sum", mySum(2, 3, 4, 51, 6))
	fmt.Println("Total Sum", mySum(2, 3, 43, 5, 6))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
