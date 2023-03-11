package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	// range keeps looping over the channel until it is complete
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send only channel
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// close the channel otherwise the for loop in main function
	// keeps waiting
	close(c)
}
