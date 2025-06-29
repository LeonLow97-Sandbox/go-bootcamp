package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// will keep ranging over the previous read-only channel until it is closed
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	// input
	nums := []int{2, 3, 4, 7, 1, 12}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2: pass output from stage 1 into another stage
	finalChannel := sq(dataChannel)

	// stage 3: output result of entire pipeline
	// will continuously range over the output channel until it is closed
	// output is in order because we use unbuffered channel and communication
	// is synchronous
	for n := range finalChannel {
		fmt.Println(n)
	}
}
