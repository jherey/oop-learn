package payment

import "fmt"

// Cash data structure
type Cash struct{}

// CreateCashAccount constructor function to create a cash account
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment processes cash payment
func (c Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}
