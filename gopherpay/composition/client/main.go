package main

import "fmt"

// Account data structure
type Account struct{}

// AvailableFunds returns the accounts available funds
func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}

// ProcessPayment processes a payment
func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

// CreditAccount data type
type CreditAccount struct {
	Account
}

func main() {
	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)
}
