package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// copy with serialization (encoders and decoders)
func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Person{}    // prepare memory for Person
	_ = d.Decode(&result) // decode into an object from the buffer
    return &result
}

func main() {
	john := Person{
		"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Darrel", "Alfred"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Terizla")

	fmt.Println("John", john, john.Address)
	fmt.Println("Jane", jane, jane.Address)
}
