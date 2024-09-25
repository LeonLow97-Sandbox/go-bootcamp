package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// NewPerson is a factory function that creates and returns a new Person
func NewPerson(name string, age int) *Person {
	// can perform some validation
	if age < 16 {
		// ...
	}

	// Create a new Person instance and initialize its fields
	return &Person{name, age, 2}
}

func main() {
	// Use the factory function NewPerson to create a Person object
	p := NewPerson("John", 33)
	fmt.Println(p)
}
