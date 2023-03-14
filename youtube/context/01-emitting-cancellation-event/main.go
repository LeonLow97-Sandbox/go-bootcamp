package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	// Create a new context
	ctx := context.Background()

	// Create a new context, with its cancellation function
	// from its original context
	ctx, cancel := context.WithCancel(ctx)

	// Run 2 operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	} ()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}

func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("Failed. Took too long.")
}

func operation2(ctx context.Context) {
	// use a similar pattern to the HTTP server
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Done!")
	case <-ctx.Done():
		fmt.Println("Halted Operation2!")
	}
}
