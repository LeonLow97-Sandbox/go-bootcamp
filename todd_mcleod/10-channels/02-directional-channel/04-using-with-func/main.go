package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("about to exit")
}

// send only channel
func foo(c chan<- int) {
	c <- 42
}

// receive only channel
func bar(c <-chan int) {
	fmt.Println(<-c)
}
