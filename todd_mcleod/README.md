# Golang Help

- <a href="https://go.dev/ref/spec" target="_blank">Golang Language Specification</a>
- <a href="https://gobyexample.com/" target="_blank" >Go by Example</a>
- <a href="https://go.dev/play/" target="_blank">Go Playground</a>
- <a href="https://pkg.go.dev/std" target="_blank">Go Package Documentation</a>
- <a href="https://godoc.org">GoDoc</a>

# Content

<a href="#function">Functions</a> | <a href="#structs">Structs</a>

# Variables, Values and Type

# Different Types

# Conditionals

# Composite Data Types

<h1 id="structs">Structs</h1>

### What is a struct?

- A struct is a data structure that **compose together values of different types**.

```
type person struct{}
```

- Empty struct

```
type person struct {
	first string
	last  string
	age int
}
```

- Struct fields must be unique. Cannot have another `first string`.
- Must have an identifier and type.

### Embedded Struct

```
type secretAgent struct {
	person	// another struct
	ltk bool
}
```

- Anonymous field: a field declared with a type but no explicit field name.
- A field from embedded struct is promoted to the outer level.

### Anonymous Struct

```
	p3 := struct {
		first string
		last  string
		age int
	}{
		first: "Tuck Yern",
		last:  "Chan",
		age: 25,
	}
```

<h1 id="function">Functions</h1>

### Syntax

- `func (r receiver) identifier(parameters) (return(s)) {...}`
- **Difference between parameters and arguments**: we define our func with parameters (if any) and we call our func and pass in arguments.

- Example Syntax:

```
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := true
	return a, b
}
```

### Variadic Parameter

```
func foo(x ...int) int {
```

- Takes in an unlimited number of int as the parameter.

```
z := foo()
```

- **Variadic Parameter**: Passing in 0 or more arguments. Must be the final parameter (...T)
- `func foo(s string, x ...int)` allowed. `z := foo("Leon")`
- `func foo(x ...int, s string)` not allowed.

### Unfurling a Slice

```
	yi := []int{2,3,4,5,6,7,8,9}
	y := foo(yi...)
	fmt.Println(y)
```

- Using the `...` notation, can unpack a slice into the integers.

### defer

- `defer` is a Keyword.
- `defer`: defer the execution of a function until whereever it is being called comes to an end.

```
func main() {
	defer foo()
	bar()
} // foo() is ran right here before the function ends

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
```

- Deferring the function `foo()` in `main()`

### Receiver

```
func (s secretAgent) speak()
```

- Any value of type secretAgent has access to this method.
- Attaches the function `speak()` to the type `secretAgent`

### Interface

```
type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}
```

- Anything that has the method `speak()` is also of type `human`.

```
func bar(h human) {
	switch h.(type) {  // this is an assert; asserting, "h is of this type"
	case person:
		fmt.Println("I was passed into bar!!!", h.(person).first)
	case secretAgent :
		fmt.Println("I was passed into bar!!!", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}
```

- **Assertion** using `switch` statement. Asserting the type.

### Anonymous Function

- Anonymous Function with no parameters

```
func main() {

	// Anonymous Function
	func() {
		fmt.Println("Anonymous function ran")
	}()

}
```

- Anonymous Function with parameters

```
func (x int) {
	fmt.Println("The meaning of life:", x)
} (42)
```

### Func Expression

- Assigning function to variables

```
func main() {
	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	y := func(x int) {
		fmt.Println("The year I was born in is", x)
	}

	y(2003)
}
```

### Returning a Function

- Return a string

```
func main() {
	s1 := foo()
	fmt.Println(s1)
}

func foo() string {
	s := "Hello World!"
	return s
}
```

- Returns a function that returns an int

```
func bar() func() int {
	return func() int {
		return 451
	}
}

func main() {
	x := bar()
	fmt.Println(x())  // calling func() int
	fmt.Printf("%T\n", x())
}
```

- Can also run `bar()` entirely
  `fmt.Println(bar()())`

### Function Callback

- Passing a function as an argument to another function.

```
func even(f func(xi ...int) int, vi ...int) int {
	// Finding the even number
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func main() {
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)
}
```

### Closure

- Using a block or scope to enclose some variables.

### Recursive

- **Recursion**: When a function calls itself.

```
func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)
}

func factorial (n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}
```

# Additional Code

|         Code         |               Description                |
| :------------------: | :--------------------------------------: |
| `func foo(x ...int)` | Passing in an unlimited number of `int`. |
|        `s...`        |            Unfurling a Slice.            |

# Pointers

### What are pointers?

- Storing values in Golang creates a memory address.

|  Code  |                               Description                               |
| :----: | :---------------------------------------------------------------------: |
|  `&a`  |               Gives you the **address**. `"Pointer to a"`               |
|  `*a`  | Gives you the **value** stored at an address. "Deferencing an address". |
| `*int` |               Pointer to the type `int` stored in memory.               |

### When to use pointers?

- When you have a large chunk of data, and you don't want to pass the huge chunk around, just use the address where that value is stored.

### Method Sets

- Method sets determine what methods attach to a TYPE.

|     Method Sets      |                     Description                      |
| :------------------: | :--------------------------------------------------: |
| Non-Pointer Receiver | Works with values that are POINTERS or NON-POINTERS. |
|   Pointer Receiver   |      ONLY works with values that are POINTERS.       |

| Receivers |  Values   |
| :-------: | :-------: |
|  `(t T)`  | T and \*T |
| `(t *T)`  |    \*T    |

# Application

### JSON Marshal

- `func Marshal(v any) ([]byte, error)`
- Returns a slice of bytes and error
- Taking data and turning it into JSON

```
	people := []person{p1, p2}

	fmt.Println(people)

	// Marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	// Convert byte slice to string
	fmt.Println(string(bs))
```

### JSON UnMarshal

- `func Unmarshal(data []byte, v any) error`
- Unmarshal parses the JSON-encoded data and stores the result in the value pointer to by v (data structure).
- <a href="https://mholt.github.io/json-to-go/">Convert JSON to Go</a>
- Receiving JSON data and putting it into GoLang Data Structure.

### Writer Interface

- Encode: Send out as JSON
- Decode: Bring in JSON
- Write it to the output source
- Can be a file or web connection, send it to client.

- type `Writer`

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- In the example, we used Fprintln which takes in a writer interface <br>
  `func Fprintln(w io.Writer, a ...any) (n int, err error)` <br>
  `n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")`
- `os.Stdout` is of type Writer
- func `WriteString` <br>
  `func WriteString(w Writer, s string) (n int, err error)`

### Sort

```
func main() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
}
```

- Sorting a slice of int in ascending order

### Sort Custom

```
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	// define a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	// type conversion: `people` now is type `ByAge`
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
```

### bcrypt

- To install bcrypt:

  - `go env` to check GO111MODULE value.
  - If not set to `auto`, run `go env -w GO111MODULE=auto`
  - `go get golang.org/x/crypto/bcrypt`

- <a href="https://pkg.go.dev/golang.org/x/crypto/bcrypt">Bcrypt Link</a>

- For encrypting and storing passwords, a hashing function. <br>
  `func GenerateFromPassword(password []byte, cost int) ([]byte, error)`

- CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent
- Returns nil on success, or an error on failure.
  `func CompareHashAndPassword(hashedPassword, password []byte) error`

- Cost

```
const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)
```

# Concurrency

### Concurrency vs Parallelism

- Concurrency: can have multiple threads executing code. If 1 thread blocks, another one is picked up and worked on.
- Parallelism: multiple threads executed at the exact same time. Requires multiple CPUs.

### WaitGroup

- `WaitGroup` waits for a collection of goroutines to finish.
- The main goroutine calls `Add` to set the number of goroutines to wait for.
- Then each of the goroutines runs and calls `Done` when finished.
- At the same time, `Wait` can be used to block until all goroutines have finished.
- <a href="https://godoc.org/sync">Go Doc WaitGroup</a>

### Method Sets Revisited

- A type may have a method set associated with it.
- The method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type.

### Race Condition

- Race conditions are the outcomes of 2 different concurrent contexts reading and writing to the same shared data at the same time, resulting in an unexpected output.
- In Golang, 2 concurrent goroutines that access the same variable concurrently will produce a data race in the program.

### `runtime.Gosched()`

- Gosched yields the processor, allowing other goroutines to run.
- It does not suspend the current goroutine, so execution resumes automatically.

### Mutex

- To prevent a race condition, so that multiple go routines can access that same code at the same time.
- `Lock`: If the lock is already in use, the calling goroutine blocks until the mutex is available.
- `Unlock`: It is allowed for one goroutine to lock a Mutex and then arrange for another goroutine to unlock it.

### Atomic

- Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.
- These function require great care to be used correctly.

- The add operation, implemented by the `AddT` functions, is the atomic equivalent of:

```
*addr += delta
return *addr
```

- `func AddInt64(addr *int64, delta int64) (new int64)`
- `AddInt64` automatically adds delta to \*addr and returns the new value.

# Channels

### Channels

|          Channels           |                Code                 |
| :-------------------------: | :---------------------------------: |
|      Making a channel       |        `c := make(chan int)`        |
| Putting values on a channel |              `c <- 42`              |
| Taking values off a channel |               `<- c`                |
|  Making a buffered channel  | `c := make(chan int, <buffer no.>`) |

### Understanding Channels

- `Concurrency`: have multiple goroutines running at the same time.

- The following code does not work.
- This is because "channels block"

```
func main() {
	// creates a channel
	c := make(chan int)

	// blocks the channel in this single goroutine main
	c <- 42

	fmt.Println(<-c)
}
```

- The following code works.
- 2 goroutines are created, one of which receives value in the channel
- At the end of the code execution, the value was taken out of the channel.

```
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
```

- Another method (buffer) of a working channel

```
func main() {
	// buffer channel: allows value to sit in there
	// 1 is the value in this case
	c := make(chan int, 1)

	// 42 gets put into the channel because there is a buffer of 1
	// no longer blocking the channel
	c <- 42

	fmt.Println(<-c)
}
```

- Unsuccessful buffer
- Added another value to be put into the channel but the channel only has a buffer of 1

```
func main() {
	// buffer channel: allows value to sit in there
	// 1 is the value in this case
	c := make(chan int, 1)

	// 42 gets put into the channel because there is a buffer of 1
	// no longer blocking the channel
	c <- 42
	c <- 43

	fmt.Println(<-c)
}
```

- Successful Buffer

```
func main() {
	// buffer channel: allows value to sit in there
	// 1 is the value in this case
	c := make(chan int, 1)

	// 42 gets put into the channel because there is a buffer of 1
	// no longer blocking the channel
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
```

### Directional Channels

- Read channels from left to right

| Directional Channel |           Code            |                       Description                        |
| :-----------------: | :-----------------------: | :------------------------------------------------------: |
|       Receive       | `cr := make(<-chan int)`  | data only comes out (can read from, but cannot write to) |
|        Send         | `cs := make(chan <- int)` |  data only goes in (cannot read from, but can write to)  |

- Error in receiving from channel

```
func main() {
	c := make(chan <- int, 2)  // channel is going to be sent an int

	c <- 42
	c <- 43

	fmt.Println(<-c) // error: trying to receive from the channel too
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
```

- Receive only channel

```
func main() {
	c := make(<- chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
```

### Ranging over a channel

- `close()` waits for the channel to complete. 
```
func main() {

	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send only channel
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
```

### `select` statement

```
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
```

# Error Handling

|Error|Description|
|:-:|:-:|
|`fmt.Println()`|Print out to os.Stdout|
|`log.Println()`|Can choose the location to display the error|
|`log.Fatalln()`|`os.Exit()`. Shut down. Fatal function call os.Exit(1) after writing the log message|
|`log.Panicln()`|Deferred functions run. Can use "recover" from an error.|
|`panic()`|Another type of error.|


- Fatal functions call ox.Exit(1) after writing the log message...
- Fatalln is equivalent to Println() followed by a call to os.Exit(1)

- Panicln is equivalent to Println followed by a call to panic()
- Fatalln is equivalent to Println() followed by a call to os.Exit(1).

### Log Error

- Logging errors might be better to save the output of errors in a file. 

```Golang
log.SetOutput(<Writer Interface>)
```

### Defer Panic and Recover

[Defer, Panic and Recover](https://go.dev/blog/defer-panic-and-recover)
- Recover is only useful inside deferred functions.

### Error Type

```
type error interface {
	Error() string
}
```

# Testing & Benchmarking

### Introduction

- Tests must
	- be in a file that ends with `_test.go` 
	- put the file in the same package as the one being tested
	- be in a func with a signature `func TestXx(*testing.T)`

- Run a test
	- `go test`

- Deal with a test failure
	- use the Error, Fail or related methods to signal failure.

[Golang Testing](https://www.golang-book.com/books/intro/12)

### Run all folder beneath the current directory

`go test ./...`

### Example Code

```
func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}
```

### Golint

- Golint is a linter for Go source code.
- Golint differs from gofmt. Gofmt reformats Go source code, whereas golint prints out style mistakes.
- Golint differs from govet. Govet is concerned with correctness, whereas golint is concerned with coding style.

- `go get -u golang.org/x/lint/golint`
- `go list -f {{.Target}} golang.org/x/lint/golint`



