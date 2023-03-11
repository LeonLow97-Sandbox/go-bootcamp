package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {
	// Assignment 1: Create a value and assign it to a variable
	// Print the address of that value
	x := 67
	fmt.Println(&x)

	p := person{"Leon", "Low", 25}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	p.first = "Jie Wei"
	// (*p).first = "Jie Wei"
	p.last = "Low"
	fmt.Println(p.first, p.last)
}