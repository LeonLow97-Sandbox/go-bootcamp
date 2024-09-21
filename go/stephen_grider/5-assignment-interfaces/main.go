package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	mySquare := square{sideLength: 4}
	myTriangle := triangle{base: 8, height: 6}

	// printArea(mySquare)
	// printArea(myTriangle)
	printArea(mySquare)
	printArea(myTriangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
