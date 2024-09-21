package main

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {}
func (m MultiFunctionPrinter) Fax(d Document)   {}
func (m MultiFunctionPrinter) Scan(d Document)  {}

type OldFashionedPrinter struct{}

func (m OldFashionedPrinter) Print(d Document) {}

// Deprecated: ...
func (m OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (m OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {}
func (p Photocopier) Scan(d Document)  {}

// compose interfaces
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {}
