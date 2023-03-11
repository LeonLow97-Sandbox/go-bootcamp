package main

import "fmt"

func main() {
	bar("Leon")
	s1 := woo("Moneypenny")
	fmt.Println(s1)

	// returns 2 variables
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// Everything in Go is PASS BY VALUE
// Basic Syntax of Function
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := true
	return a, b
}