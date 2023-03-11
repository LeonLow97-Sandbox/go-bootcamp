package main

import "fmt"

func main() {
	// Build and use an anonymous function
	func() {
		fmt.Println("Hello World!")
	}()

	func(x int) {
		fmt.Println(x)
	}(10)

	func () {
		for i := 0; i <= 20; i ++ {
			fmt.Println(i)
		}
	} ()
}