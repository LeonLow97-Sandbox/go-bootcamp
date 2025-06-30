package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// defining a custom struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// reference to struct in memory
	jim.updateName("Jimmy")
	jim.print()
}

// receiver type *person (type of pointer that points at person)
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// receiver function with 'person' type
func (p person) print() {
	fmt.Printf("%+v", p)
}
