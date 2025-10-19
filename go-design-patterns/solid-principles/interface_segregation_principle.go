package main

// ISP - Interface Segregation Principle
// Clients should not be forced to depend on interfaces they do not use.

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

func (o OldFashionedPrinter) Print(d Document) {}
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Better approach: split into several interfaces
type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Scan(d Document)  {}
func (p Photocopier) Print(d Document) {}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combination + decorator
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

func ISP() {
	// Violate ISP
	var printer Machine = OldFashionedPrinter{}
	printer.Print(Document{})
	// printer.Scan(Document{}) // panic: operation not supported

	// Adhere to ISP
	var myPrinter Printer = MyPrinter{}
	myPrinter.Print(Document{})

	var photocopier Scanner = Photocopier{}
	photocopier.Scan(Document{})

	var mfd MultiFunctionDevice = MultiFunctionMachine{
		printer: myPrinter,
		scanner: photocopier,
	}
	mfd.Print(Document{})
	mfd.Scan(Document{})
}
