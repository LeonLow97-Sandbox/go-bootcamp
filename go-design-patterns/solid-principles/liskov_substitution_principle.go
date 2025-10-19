package main

import "fmt"

// LSP - Liskov Substitution Principle
// Not applicable to Go because Go does not have traditional class inheritance.

type Sized interface {
	GetWidth() int
	SetWidth(w int)
	GetHeight() int
	SetHeight(h int)
}

type Rectangle struct {
	width, height int
}
func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(w int) {
	r.width = w
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(h int) {
	r.height = h
}

type Square struct {
	Rectangle
}
func NewSquare(size int) *Square {
	sq := &Square{}
	sq.width = size
	sq.height = size
	return sq
}
// Breaks LSP because SetWidth and SetHeight do not behave as expected
// SetWidth should not change height and vice versa
func (s *Square) SetWidth(w int) {
	s.width = w
	s.height = w
}
func (s *Square) SetHeight(h int) {
	s.height = h
	s.width = h
}

func UseIt(s Sized) {
	w := s.GetWidth()
	s.SetHeight(10)
	expectedArea := 10 * w
	actualArea := s.GetWidth() * s.GetHeight()
	fmt.Printf("Expected area: %d, got %d\n", expectedArea, actualArea)
}

func LSP() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
