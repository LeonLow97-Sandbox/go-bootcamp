package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	fmt.Println("Hello I am number one")
	wg.Done()
}

func bar() {
	fmt.Println("Hello I am number two")
	wg.Done()
}