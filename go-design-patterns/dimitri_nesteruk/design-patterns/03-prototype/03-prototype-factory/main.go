package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

// copy with serialization (encoders and decoders)
func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}  // prepare memory for Employee
	_ = d.Decode(&result) // decode into an object from the buffer
	return &result
}

// 2 prototypes, we want to make copies of these and customize them
var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London"},
}
var auxOffice = Employee{
	"", Address{0, "66 West Dr", "London"},
}

// newEmployee takes a prototype Employee and customizes it by modifying the name and suite number.
func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()

	// performed customization here on the prototype
	result.Name = name
	result.Office.Suite = suite

	return result
}

// factory functions
// using the newEmployee function to create new instances based on the prototypes.
func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}
func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
