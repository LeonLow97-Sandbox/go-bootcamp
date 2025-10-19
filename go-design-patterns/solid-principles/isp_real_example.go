package main

import "fmt"

// ISP demonstrates the Interface Segregation Principle.
// Clients should not be forced to depend on interfaces they do not use.

// Violate ISP (Fat interface)
type PaymentGateway interface {
	Authorize(amount float64) error
	Capture(transactionID string) error
	Refund(transactionID string) error
	GenerateInvoice(transactionID string) string
}
type StripeGateway struct{}

func (s *StripeGateway) Authorize(amount float64) error {
	return nil
}
func (s *StripeGateway) Capture(transactionID string) error {
	return nil
}
func (s *StripeGateway) Refund(transactionID string) error {
	return nil
}
func (s *StripeGateway) GenerateInvoice(transactionID string) string {
	fmt.Println("Generating invoice for stripe transaction:", transactionID)
	return "Stripe Invoice"
}

type InternalGateway struct{}

func (p *InternalGateway) Authorize(amount float64) error {
	return nil
}

// InternalGateway does not need to implement Capture, Refund, or GenerateInvoice
// but is forced to do so due to the fat interface
func (p *InternalGateway) Capture(transactionID string) error {
	panic("operation not supported")
}
func (p *InternalGateway) Refund(transactionID string) error {
	panic("operation not supported")
}
func (p *InternalGateway) GenerateInvoice(transactionID string) string {
	panic("operation not supported")
}

// ISP: segregated fat interface into smaller interfaces
type Authorizer interface {
	Authorize(amount float64) error
}
type Capturer interface {
	Capture(transactionID string) error
}
type Refunder interface {
	Refund(transactionID string) error
}
type InvoiceGenerator interface {
	GenerateInvoice(transactionID string) string
}

type StripeGatewayV2 struct{}

func (s *StripeGatewayV2) Authorize(amount float64) error {
	return nil
}
func (s *StripeGatewayV2) Capture(transactionID string) error {
	return nil
}
func (s *StripeGatewayV2) Refund(transactionID string) error {
	return nil
}
func (s *StripeGatewayV2) GenerateInvoice(transactionID string) string {
	fmt.Println("Generating invoice for stripe transaction:", transactionID)
	return "Stripe Invoice"
}

type InternalGatewayV2 struct{}

// InternalGatewayV2 only implements Authorizer interface
// No need to implement other methods
func (p *InternalGatewayV2) Authorize(amount float64) error {
	return nil
}

func ISPExample() {
	// Violates ISP
	var stripeGateway PaymentGateway = &StripeGateway{}
	stripeGateway.Authorize(100)
	stripeGateway.Capture("txn_123")
	stripeGateway.Refund("txn_123")
	stripeGateway.GenerateInvoice("txn_123")

	var internalGateway PaymentGateway = &InternalGateway{}
	internalGateway.Authorize(200)

	// ISP
	stripeGatewayV2 := &StripeGatewayV2{}
	stripeGatewayV2.Authorize(100)
	stripeGatewayV2.Capture("txn_123")
	stripeGatewayV2.Refund("txn_123")
	stripeGatewayV2.GenerateInvoice("txn_123")

	internalGatewayV2 := &InternalGatewayV2{}
	internalGatewayV2.Authorize(200)
}
