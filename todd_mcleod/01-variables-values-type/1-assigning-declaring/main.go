package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 124
// declare & assign = initialization
var y = 100 + 24

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE INT
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// type: false for booleans, 0 for numeric types, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int
var a = "Hello Jie Wei"   // type string
var b = `Hello Leon`   // type String with backticks

var zeroBool bool

func main() {
	assign()
	foo()
	printType()

	fmt.Println(z)
	fmt.Println(zeroBool)

	c := 60
	// other types of format string
	fmt.Printf("%b\n", c)
	fmt.Printf("%x\n", c)
	fmt.Printf("%#x\n", c)
	fmt.Printf("%#x\t%b\t%x\t\n", c, c, c)

	fmt.Print("Jie")
	fmt.Print("Wei\n")

	// string printing
	s := fmt.Sprintf("%#x\t%b\t%x\t\n", c, c, c)
	fmt.Println(s)

	name := "Leon Low"
	fmt.Printf("%v", name)
}



func assign() {
	// short declaration operatpr
	// declare a variable and assign a value (of a certain TYPE)
	x := 42
	fmt.Println(x)
	// reassign
	x = 99
	fmt.Println(x)

	// "var"
	fmt.Println(y)
}

func foo () {
	fmt.Println("again", y)
}

func printType() {
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}