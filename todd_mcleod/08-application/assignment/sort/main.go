package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 12, 2, 567, 32, 71, 1256, 921}
	xs := []string{"random", "rainbow, starbucks", "boeing", "tesla", "visa", "apple", "microsoft", "javascript"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}