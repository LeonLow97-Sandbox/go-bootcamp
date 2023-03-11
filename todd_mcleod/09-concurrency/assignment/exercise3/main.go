package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		v := counter
		runtime.Gosched()
		v++
		counter = v
		wg.Done()
	}()

	go func() {
		v := counter
		runtime.Gosched()
		v++
		counter = v
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(counter)
}