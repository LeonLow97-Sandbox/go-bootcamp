package main

import "fmt"

func main() {
	c := make(chan<- int, 2) // receives int

	c <- 42
	c <- 43

	fmt.Println(c) // error: trying to receive from the channel too
	fmt.Println(c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
