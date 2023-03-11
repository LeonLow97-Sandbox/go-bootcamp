package main

import (
	"fmt"
)

func main() {
	// c := make(chan int, 1)
	c := make(chan int)

	// c <- 42

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
