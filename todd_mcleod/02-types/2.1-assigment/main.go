package main

import "fmt"

const (
	a = 42
	b int = 42
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

const (
	_ = iota + 2017
	year1
	year2
	year3
	year4
)

func exercise6() {
	// Using iota, create 4 constants for the last 4 years. Print the constant values
	fmt.Println(year1, year2, year3, year4)
}

func exercise5() {
	// Create a variable of type string using a raw string literal and print it
	a := `here is something
	as
	a
	raw
	string literal "HAHA" with quotations`
	fmt.Println(a)
}

func exercise4() {
	// bits shifting
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}

func exercise3() {
	// Creating Typed and Untyped constants
	fmt.Println(a, b)
}

func exercise2() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Println(a, b, c, d, e, f)
}

func exercise1() {
	// Write a program that prints a number in decimal, binary and hex
	x := 42

	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)
}