package main

import "fmt"

// CreditAccount data structure
type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment...")
}

// CreateCreditAccount constructor function to create a credit card
func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}

	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}

func main() {
	chargeCh := make(chan float32)
	CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(&a)
}
