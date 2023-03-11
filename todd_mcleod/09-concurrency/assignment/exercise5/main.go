package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
		fmt.Println("Counter\t", atomic.LoadInt64((&counter)))
		wg.Done()
	}()

	go func() {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
		fmt.Println("Counter\t", atomic.LoadInt64((&counter)))
		wg.Done()
	}()

	fmt.Println("GoRoutine", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println(counter)
}