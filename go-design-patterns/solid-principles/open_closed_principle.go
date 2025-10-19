package main

import "fmt"

// OCP - Open/Closed Principle
// Software entities (classes, modules, functions, etc.) should be open for extension but closed for modification.
// Specification pattern can be used to achieve OCP.

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

// Anti-Pattern: Violates OCP
// Every time we want to add a new filtering criteria, we have to modify the Filter class.
// func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
// 	result := make([]*Product, 0)
// 	for i, v := range products {
// 		if v.color == color {
// 			result = append(result, &products[i])
// 		}
// 	}
// 	return result
// }
// func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
// 	result := make([]*Product, 0)
// 	for i, v := range products {
// 		if v.size == size {
// 			result = append(result, &products[i])
// 		}
// 	}
// 	return result
// }
// func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
// 	result := make([]*Product, 0)
// 	for i, v := range products {
// 		if v.size == size && v.color == color {
// 			result = append(result, &products[i])
// 		}
// 	}
// 	return result
// }

// Specification Pattern to adhere to OCP
// Extend the filtering capabilities by creating new specifications without modifying existing code.
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type BetterFilter struct{} // will not be modified
func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
// Composite Specification design pattern can also be implemented to combine multiple specifications (AND, OR, NOT)
type AndSpecification struct {
	first, second Specification
}
func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

func OCP() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// Anti-Pattern usage
	// filter := Filter{}
	// greenProducts := filter.FilterByColor(products, green)
	// for _, v := range greenProducts {
	// 	fmt.Println(" - ", v.name, "is green")
	// }

	// OCP usage
	// open for extension, closed for modification, not modifying specification
	betterFilter := BetterFilter{}
	greenSpec := ColorSpecification{green}
	greenProducts := betterFilter.Filter(products, greenSpec)
	for _, v := range greenProducts {
		fmt.Println(" - ", v.name, "is green")
	}

	largeSpec := SizeSpecification{large}
	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	largeGreenProducts := betterFilter.Filter(products, largeGreenSpec)
	for _, v := range largeGreenProducts {
		fmt.Println(" - ", v.name, "is large and green")
	}
}
