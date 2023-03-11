package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		// checking if there is anything in c
		case v := <-c:
			fmt.Println(v)
		// if there is something in q, return
		case <-q:
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
