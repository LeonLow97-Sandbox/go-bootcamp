package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:

			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			// reading value from `stream`, writing value to `taken`
			case taken <- <-stream:

			}
		}
	}()

	return taken
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	// very slow operation, slow pipeline stage (intentionally creating a slow stage)
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}

func pipelinesAndGenerator() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randomNumFetcher := func() int {
		// function that returns a random integer no greater than 500M
		return rand.Intn(500000000)
	}

	randIntStream := repeatFunc(done, randomNumFetcher)

	// Fan Out
	CPUCount := runtime.NumCPU()
	fmt.Println("Number of CPUs:", CPUCount)            // equals to number of go routines we use for fan out
	primeFinderChannels := make([]<-chan int, CPUCount) // fan out, creating multiple channels
	for i := 0; i < CPUCount; i++ {
		// For each CPU, we create a channel for primeFinder (fan out)
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// Fan In
	fannedInStream := fanIn(done, primeFinderChannels...)
	for random := range take(done, fannedInStream, 10) {
		fmt.Println(random)
	}

	fmt.Println(time.Since(start))
}
