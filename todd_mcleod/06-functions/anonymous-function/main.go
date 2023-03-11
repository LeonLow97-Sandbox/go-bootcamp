package main

import "fmt"

func main() {

	// Anonymous Function
	func() {
		fmt.Println("Anonymous function ran")
	}()
	
	// Anonymous Function with parameters
	func (x int) {
		fmt.Println("The meaning of life:", x)
	} (42)
	
}