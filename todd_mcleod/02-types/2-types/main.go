package main

import "fmt"

// boolean type
var x bool

func main() {
	BooleanType()
	NumericType()
	StringType()
}

func StringType() {
	s := "Hello Jie Wei!"
	fmt.Println(s)
	fmt.Printf("%T\t%v\n", s, s)   // strings are immutable

	// Strings are made out of a slice of bytes
	// Converting string to a slice of bytes
	// each letter represneted as decimal numbers
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// print out UTF8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	// print new line
	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}
}

func NumericType() {
	x := 42
	y := 42.12314
	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var a int8 = 127
	fmt.Println(a)

	var b float64 = 54.243234
	fmt.Println(b)

	// var a int
	// a = 2.35 // error
}

func BooleanType() {
	a := 7
	b := 42

	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(x)
	x = true
	fmt.Println(x) 
}