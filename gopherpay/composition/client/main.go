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

// CreditAccountType data type
type CreditAccountType struct {
	Account
}

func main() {
	ca := &CreditAccountType{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)
}
