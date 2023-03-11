package main

import "fmt"

// Create a user defined struct with the identifier "person"
// Fields : `first`, `last`, `age`
type person struct {
	first string
	last string
	age int
}

func main() {
	// Create a value of type person
	p := person {
		first: "Leon",
		last: "Low",
		age: 25,
	}

	p.speak()
}

// Attach a mathod to type person with the identifier "speak"
// The method should have the person say their name
func (p person) speak() {
	fmt.Println("Hi my name is", p.first, p.last, "and my age is", p.age )
}