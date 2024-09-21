package main

import "fmt"

// Dependency Inversion Principle (it is not exactly dependency injection)
// High-Level modules should not depend on Low-Level modules
// Both should depend on abstractions (interfaces in Go)

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// other attributes
}

// model relationship between 2 people
type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module (e.g., in our case, like a storage, could also be a database)
// storing the relationships between people
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// break DIP, hlm depending on llm
	// relationships Relationships

	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	// relations := r.relationships.relations // relationships is a low-level module and our high-level module is relying on it
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		fmt.Println("John has a child called", rel.to.name)
	// 	}
	// }

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// llm
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
