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

	// mutex
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// sleeps for a duration
			// time.Sleep(time.Second)
			// runtime.Gosched() yields the processor, allowing other goroutines to run.
			// does not suspend the current goroutine, so other execution resumes automatically.
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutine", runtime.NumGoroutine())
	fmt.Println("count", counter)
}