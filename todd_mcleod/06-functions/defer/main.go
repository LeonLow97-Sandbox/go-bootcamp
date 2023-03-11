package main

import "fmt"

func main() {
	defer foo()
	bar()
} // foo() is ran right here before the function ends

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}