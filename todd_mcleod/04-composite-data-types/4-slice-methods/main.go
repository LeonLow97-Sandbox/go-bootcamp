package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 56}

	// Append the elements to the end of the slice and return the result
	// append(slice, int elements...)
	x = append(x, 6)
	fmt.Println(x)

	// Append a slice to a slice
	y := []int{7123,128,92}
	x = append(x, y...)  // spread operator of slice y (variadic parameter)
	fmt.Println(x)

	// Delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// make (built-in function)
	// takes a slice, length, capacity
	a := make([]int, 10, 10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 2423)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a)) // capacity changes to 20 (created another similar capacity because no space)

	// Multi-dimensional Slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	xp := [][]string{jb, mp}  // slice of a slice of string
	fmt.Println(xp)

	fmt.Println(xp[1][1])
}