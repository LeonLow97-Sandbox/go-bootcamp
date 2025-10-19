package main

import "fmt"

// Violates ISP
// Low-Level Module
type VisaGateway struct{}

func (s *VisaGateway) ProcessPayment(amount float64) {
	fmt.Println("Processing visa payment of", amount)
}

// High-Level Module
// violates DIP because it depends on low-level module VisaGateway
// If we change to PayPal later, we need to modify this class
// Violates both OCP (hard to extend to other gateways) and DIP
type PaymentService struct{}

func (s *PaymentService) Pay(amount float64) {
	visa := VisaGateway{}
	visa.ProcessPayment(amount)
}

// Adhere to ISP
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type PaymentServiceDIP struct {
	processor PaymentProcessor // HLM depends on abstraction
}
func (s *PaymentServiceDIP) Pay(amount float64) {
	s.processor.ProcessPayment(amount)
}

type MastercardGateway struct{}
func (s *MastercardGateway) ProcessPayment(amount float64) {
	fmt.Println("Processing mastercard payment of", amount)
}

func DIPExample() {
	// Violates ISP
	paymentService := PaymentService{}
	paymentService.Pay(100.0) // but this always pays via Visa, hard to extend

	// Adhere to ISP (easily extend to other gateways)
	// Abstraction: PaymentProcessor
	// High-Level Module: PaymentServiceDIP
	// Low-Level Modules: VisaGateway, MastercardGateway
	mastercard := &MastercardGateway{}
	paymentServiceDIP := PaymentServiceDIP{processor: mastercard}
	paymentServiceDIP.Pay(200.0) // pays via Mastercard, easy to extend

	visa := &VisaGateway{}
	paymentServiceDIP2 := PaymentServiceDIP{processor: visa}
	paymentServiceDIP2.Pay(100.0) // pays via Visa, easy to extend
}
