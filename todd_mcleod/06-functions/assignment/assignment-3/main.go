package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello World!")
}

// Use the `defer` keyword to show that a deferred func
// runs after the func containing it exits
func foo() {

	defer func(){
		fmt.Println("I will be printed last :(")
	}()

	fmt.Println("HAHA I am second!")

}