# Material Link

- [Request Context](https://www.sohamkamani.com/golang/context-cancellation-and-values/)

# Context Values

- Can use the context variable to pass around values that are common across an operation.
- More idiomatic alternative to just passing them around as variables throughout function calls.

## Example

- For example, consider an operation that has multiple function calls, with a common ID used to identify it for logging and monitoring.

---

### Example: Naive Way

- The naive way to implement this is just pass the ID around for each function call:
- This can quickly get bloated when you have more information that you want to provide.

```go
func main() {
	// create a random integer as the ID
	rand.Seed(time.Now().Unix())
	id := rand.Int()
	operation1(id)
}

func operation1(id int64) {
	// do some work
	log.Println("operation1 for id:", id, " completed")
	operation2(id)
}

func operation2(id int64) {
	// do some work
	log.Println("operation2 for id:", id, " completed")
}
```

---

### Example: with context

- Can implement the same functionality using context:
- Here, we are creating a new context variable in `main` with a key value pair associated with it.
- The value can then be used by the successive function calls to obtain contextual information.

```go
// we need to set a key that tells us where the data is stored
const keyID = "id"

func main() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	operation1(ctx)
}

func operation1(ctx context.Context) {
	// do some work

	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	operation2(ctx)
}

func operation2(ctx context.Context) {
	// do some work

	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}
```
