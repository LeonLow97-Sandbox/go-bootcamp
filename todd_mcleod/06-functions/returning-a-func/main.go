package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)  // type 'func() int'
	fmt.Println(x())  // calling func() int

	fmt.Println(bar()())
}

func foo() string {
	return "Hello World!"
}

// Returns a func that returns an int
func bar() func() int {
	return func() int {
		return 451
	}
}