package main

import "fmt"

func main() {
	a := 42

	// Address in memory where value is stored
	fmt.Println(&a)

	fmt.Printf("%T\n", a) 
	fmt.Printf("%T\n", &a) // type *int 

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	fmt.Println(*&a)

	// a is the value stored at an address
	// b is the address where this value of a is stored
	// accessing the value at that address and reassigning it
	*b = 43
	fmt.Println(a)
}