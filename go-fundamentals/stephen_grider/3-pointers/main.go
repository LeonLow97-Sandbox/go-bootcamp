package main

import "fmt"

// pass a pointer
func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
}
