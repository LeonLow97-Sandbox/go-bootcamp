# Materials Link

- [Request Context](https://www.sohamkamani.com/golang/context-cancellation-and-values/)

# Emitting a Cancellation Event

- If you have an operation that could be cancelled, you will have to emit a cancellation event.
- This can be done using the `WithCancel` function which returns a context object, and a function.

```go
ctx, fn := context.WithCancel(ctx)
```

- The function `fn` takes no argument, and does not return anything.
  - Called when you want to cancel the context.

## Example

- Consider the case of 2 dependent operations.
- Here "dependent" means if one fails, it doesn't make sense for the other to complete. 
- If we get to know early on that one of the operations failed, we would like to cancel all dependent operations.

```go
func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func main() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)

	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
```
