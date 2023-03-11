# Golang Type Assertion

## Understanding the concept of Type Assertion in Golang

- Golang is a statically typed programming language that allows for type assertion, a mechanism that enables programmers to assert the actual type of an interface value.
- In Golang, an interface value can hold a value of any type that satisfies the interface's method set.
- Type assertion allows programmers to extract the actual value of the interface value and check for the type of an interface variable.
- Type assertion can be expressed in 2 ways:
  1. Type assertion expression
  2. Type switch statement

### Type Assertion Expression

- A type assertion expression is an expression of the form `x.(T)`, where `x` is an interface value and `T` is a type.
- This expression asserts that the dynamic type of `x` is `T`, and if successful, returns the value of `x` as a value of type `T`.

```go
var i interface{} = "Hello world!"

s := i.(string)
fmt.Printf("Type of i is %T and its value is %s\n", i, s)

// Output:
// Type of i is string and its value is Hello, world!
```

- In this example, we have an interface value `i` that holds a string value "Hello, world!".
- We use a type assertion expression to extract the actual string value from the interface value and assign it to the variable `s`.

### Type Switch Statement

- A type switch is a control structure that allows programmers to switch on the dynamic type of an interface value.
- It is used to determine the actual type of the value held by the interface.

```go
func printType(i interface{}) {
  switch v := i.(type) {
    case int:
      fmt.Printf("Type of i is %T and its value is %d\n", v, v)
    case string:
        fmt.Printf("Type of i is %T and its value is %s\n", v, v)
    default:
        fmt.Printf("Type of i is %T\n", v)
  }
}

func main() {
    printType(42)
    printType("Hello, world!")
    printType(3.14159)
}

// Output:
Type of i is int and its value is 42
Type of i is string and its value is Hello, world!
Type of i is float64
```

### Cannot use Type Assertion to check non-interface variable

- Type assertion cannot be used to check the type of a non-interface variable.
- If you want to check the type of a non-interface variable, use the `reflect` package, which provides run-time reflection on types, values and structures.
- The `reflect.TypeOf()` function can be used to get the type of a variable at runtime.

```go
func main() {
	var i int = 42
  var f float64 = 3.14
	var s = "Hello"

  fmt.Println(reflect.TypeOf(i)) // int
  fmt.Println(reflect.TypeOf(f)) // float64
	fmt.Println(reflect.TypeOf(s)) // string
}
```
