package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I'm too tired")
}

// interface factory
// return interface instead of struct
// when returning interface, don't need to be pointer like *Person
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("James", 34)

	// cannot do like p.age or p.name, this is the idea of encapsulation
	// because person is lowercase, so its a private struct
	// factory only exposed the interface

	p.SayHello()

	tp := NewPerson("James", 200)
	tp.SayHello()
}
