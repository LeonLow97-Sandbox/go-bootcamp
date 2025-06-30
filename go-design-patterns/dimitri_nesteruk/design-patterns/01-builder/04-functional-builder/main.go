package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)

// PersonBuilder maintains a list of actions (personMod functions) that will be applied
// sequentially to construct a Person
type PersonBuilder struct {
	actions []personMod
}

// Called method adds an action to set the name field of the Person. It appends a closure to b.actions
// that modifies p.name with the provided name.
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

// Build executes all the actions stored in b.actions to construct and return
// a fully initialized Person object (p)
func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dimitri").WorksAsA("Developer").Build()
	fmt.Println(*p)
}
