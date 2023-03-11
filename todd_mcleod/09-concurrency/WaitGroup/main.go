package main

import (
	"fmt"
	"runtime"
	"sync"
)

// variable wait group
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)	// waiting for 1 thing
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()	// wait until all the things we have added are done
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()	// when foo is done
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}