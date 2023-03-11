package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

	n, err := fmt.Println("Hello Jie Wei", true, 25)
	fmt.Println("Number of bytes:", n)
	fmt.Println(err)
}

func foo() {
	fmt.Println("I am in function foo")
}

func bar() {
	fmt.Println("Exited the code")
}
