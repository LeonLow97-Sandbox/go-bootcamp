package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Leon",
		Last:  "Low",
		Age:   25,
	}
	p2 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	// Marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	// Convert byte slice to string
	fmt.Println(string(bs))
}

// update