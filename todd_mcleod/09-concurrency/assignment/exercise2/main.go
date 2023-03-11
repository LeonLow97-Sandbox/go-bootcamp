package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func main() {
	p1 := person {
		first: "Leon",
	}

	// does not work
	// If receiver is a pointer, the value must be a pointer to the receiver.

	// does not work
	// saySomething(p1)
	
	saySomething(&p1)
}

func (p *person) speak() {
	fmt.Println("Hello Leon!", p)
}

func saySomething(h human) {
	h.speak()
}
