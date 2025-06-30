package main

import "fmt"

// Liskov Substitution Principle (not very applicable to Go as it deals with Inheritance)

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width // violates principle, changes height here when supposed to set width only
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height // violates principle, changes width here when supposed to set height only
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of", expectedArea, ", but got", actualArea)
}

// Square2 does not inherit directly from Rectangle. It maintains its own size, which allows for
// clearer behavior and adheres to LSP
type Square2 struct {
	size int // width, height
}

// One solution to the problem
func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
