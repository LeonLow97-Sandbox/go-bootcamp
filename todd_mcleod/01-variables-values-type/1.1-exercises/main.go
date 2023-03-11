package main

import "fmt"

// Create own type with underlying value int
// "Underlying type"
type jiewei int
var x jiewei
var y int

func main() {
	exercise1()
	exercise2()

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y := int(x)
	fmt.Println("exercise 5:", y)
	fmt.Printf("%T\n", y)
}

func exercise2() {
	var x2 int
	var y2 string
	var z2 bool

	// Zero Value
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

	// exercise 3
	x2 = 42
	y2 = "James Bond"
	z2 = true

	s := fmt.Sprintf("%v\t%v\t%v", x2, y2, z2)
	fmt.Println(s) 
}

func exercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
}

