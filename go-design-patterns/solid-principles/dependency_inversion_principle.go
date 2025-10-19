package main

// Dependency Inversion Principle
// High-level modules should not depend on low-level modules.
// Both should depend on abstractions.

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from *Person
	relationship Relationship
	to *Person
}

// Low-level module (stored in database)
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}
type Relationships struct {
	relations []Info
}

// Finding all children is now put into the low-level module
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for _, rel := range r.relations {
		if rel.from.name == name && rel.relationship == Parent {
			result = append(result, rel.to)
		}
	}
	return result
}
func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{from: parent, relationship: Parent, to: child})
	r.relations = append(r.relations, Info{from: child, relationship: Child, to: parent})
}

// High-level module (business logic)
type Research struct {
	// Break DIP because it depends on low-level module
	// relationships Relationships

	// Create an abstraction to depend on instead
	browser RelationshipBrowser
}
func (r *Research) Investigate() {
	// directly accessing low-level module's data
	// This creates a tight coupling between Research and Relationships
	// If Relationships changes to use a different storage mechanism, Research will also need to change
	// relations := r.relationships.relations 
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		println("John has a child called " + rel.to.name)
	// 	}
	// }

	children := r.browser.FindAllChildrenOf("John")
	for _, c := range children {
		println("John has a child called " + c.name)
	}
}

func DIP() {
	parent := &Person{name: "John"}
	child1 := &Person{name: "Chris"}
	child2 := &Person{name: "Matt"}

	relations := Relationships{}
	relations.AddParentAndChild(parent, child1)
	relations.AddParentAndChild(parent, child2)

	r := Research{&relations}
	r.Investigate()
}
