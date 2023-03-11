package main

import (
	"fmt"
	"math"
)

// Create a type square
type square struct {
	sideLength float64
}

// Create a type circle
type circle struct {
	radius float64
}

// Create a type shape that defines an interface as
// anything that has the Area method
type shape interface {
	area() float64
}

func main() {
	// Create a value of type square and circle
	sq := square {
		sideLength: 8,
	}
	cr := circle {
		radius: 7.2,
	}
	info(sq)
	info(cr)
}

// Attach a method to each shape that calculates area and returns it
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (c circle) area() float64 {
	pi := math.Pi
	return pi * c.radius * c.radius
}

// Create a func `info` which takes type shape and then prints the area
func info(si shape) {
	fmt.Println(si.area())
}