package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutine", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// sleeps for a duration
			// time.Sleep(time.Second)
			// can run something else
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GoRoutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutine", runtime.NumGoroutine())
	fmt.Println("count", counter)
}