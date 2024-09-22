package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

// HtmlElement represents a single HTML element with a name, text content, and child elements
type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

// String generates the string representation of the HTML element
func (e *HtmlElement) String() string {
	return e.string(0)
}

// string is a recursive method that builds the HTML representation with indentation
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

// HtmlBuilder is a builder for creating HTML elements
type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

// NewHtmlBuilder initializes a new HtmlBuilder with a specified root element name
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{rootName, "", []HtmlElement{}},
	}
}

// String generates the string representation of the entire HTML structure
func (b *HtmlBuilder) String() string {
	return b.root.String()
}

// AddChild adds a child element with the specified name and text to the root element
func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{
		childName, childText, []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
}

// AddChildFluent adds a child element and returns the builder itself for method chaining
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{
		childName, childText, []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	sb.Reset()

	words := []string{"hello", "world"}
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// Using the Builder Pattern to create an HTML structure
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	// Using Fluent Interface with the Builder Pattern
	bf := NewHtmlBuilder("ul")
	bf.AddChildFluent("li", "hello").AddChildFluent("li", "world")
	fmt.Println(b.String())
}
