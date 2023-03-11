package main

import "fmt"

func main() {
	// c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()

	// fmt.Println(<-c)

	// buffer channel: allows value to sit in there
	// 1 is the value in this case
	c := make(chan int, 1)

	// 42 gets put into the channel because there is a buffer of 1
	// no longer blocking the channel
	c <- 42

	fmt.Println(<-c)
}
