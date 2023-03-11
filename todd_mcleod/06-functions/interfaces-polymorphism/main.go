package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// Interface
// Anything that has this method speak() is also of type human.
type human interface {
	speak()
}

type hotdog int

func bar(h human) {
	switch h.(type) {  // this is an assert; asserting, "h is of this type"
	case person:
		fmt.Println("I was passed into bar!!!", h.(person).first)
	case secretAgent :
		fmt.Println("I was passed into bar!!!", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Leon",
			last:  "Low",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Fang Jen",
			last:  "Lim",
		},
		ltk: true,
	}

	p1 := person{ 
		first: "Dr.",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	// bar function takes in 'human'
	// bar takes in secretAgent and person, both have method speak()
	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}