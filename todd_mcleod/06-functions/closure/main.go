package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b()) // different variable so different "x" (different scope)
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}