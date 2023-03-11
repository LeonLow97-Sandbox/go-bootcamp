package main

import "fmt"

func main() {
	// x := type{values} // Composite Literal
	x := []int{1, 2, 3, 4, 5, 32, 4, 21}
	fmt.Println(x); fmt.Printf("%T\n", x);
	fmt.Println(x[2]) // Access element in array 
	
	// Looping over slice
	for index, value := range x {
		fmt.Println(index, value)
	}

	for i:= 0; i <= len(x)-1; i++ {
		fmt.Println(i, x[i])
	}

	// Slicing a slice (excluding the last element)
	fmt.Println(x[3:])
	fmt.Println(x[:2])
	fmt.Println(x[1:3])
}