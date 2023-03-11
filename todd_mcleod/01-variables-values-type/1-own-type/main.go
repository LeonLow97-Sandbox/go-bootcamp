package main

import "fmt"

var a int

// creating our own type
type hotdog int
var b hotdog

func main() {
	b = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)   // type: main.go

	a = 60
	// cannot assign a to b as there are different types
	// a = b // this is wrong

	// Converting b to int
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}