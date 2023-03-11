package main

import "fmt"

type person struct {
	first string
	last  string
	age int
}

type secretAgent struct {
	person	// another struct
	ltk bool
}

func main() {
	// Created a value of type person
	p1 := person{
		first: "Leon",
		last:  "Low",
		age: 25,
	}

	p2 := person{
		first: "Fang Jen",
		last:  "Lim",
		age: 24,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println("")

	// Embedded Struct
	sa1 := secretAgent {
		person: person {
			first: "Jie En",
			last: "Low",
			age: 27,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.person.first, sa1.last, sa1.age, sa1.ltk)

	fmt.Println("")

	// Anonymous Struct (Purpose: dont want messy code if you only have one struct)
	// instead of 
	// p3 := person {
	// 	first: "Tuck Yern",
	// 	last:  "Chan",
	// 	age: 25,
	// }

	// this works the same
	p3 := struct {
		first string
		last  string
		age int
	}{
		first: "Tuck Yern",
		last:  "Chan",
		age: 25,
	} 

	fmt.Println(p3)
}