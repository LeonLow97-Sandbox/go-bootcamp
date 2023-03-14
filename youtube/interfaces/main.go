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

type helper struct {
	person
	assist string
	ltk    bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (h helper) speak() {
	// different way of printing speak
	fmt.Println("I am", h.first, h.last, h.assist, " - the helper speak")
}

func bar(h human) {
	fmt.Println("I was passed into bar")
	h.speak()

	switch h.(type) {
	case helper:
		fmt.Println("Passed into bar in switch case", h.(helper).assist)
	case secretAgent:
		fmt.Println("Passed into bar in switch case", h.(secretAgent).first)
	}

}

type human interface {
	speak()
}

func main() {
	sa := secretAgent{
		person{
			"Leon", "Low",
		},
		true,
	}

	h := helper{
		person{
			"Nicholas", "Tan",
		},
		"Assisting...",
		false,
	}

	bar(sa)
	bar(h)

}
