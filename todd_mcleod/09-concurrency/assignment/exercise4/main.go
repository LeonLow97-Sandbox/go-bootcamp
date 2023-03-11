package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go func() {
		mu.Lock()
		v := counter
		
		v++
		counter = v
		mu.Unlock()
		wg.Done()
	}()

	go func() {
		mu.Lock()
		v := counter
		
		v++
		counter = v
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(counter)
}