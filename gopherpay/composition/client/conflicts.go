package main

import "fmt"

// CreditAccount credit account
type CreditAccount struct{}

// AvailableFunds returns available balance
func (c *CreditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 250
}

// CheckingAccount checking account type
type CheckingAccount struct {}

// AvailableFunds returns available balance
func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds")
	return 125
}

// HybridAccount both credit and checking accounts
type HybridAccount struct {
	CreditAccount
	CheckingAccount
}

// AvailableFunds returns available balance
func (h *HybridAccount) AvailableFunds() float32 {
	return h.CreditAccount.AvailableFunds() + h.CheckingAccount.AvailableFunds()
}

func mains() {
	ha := &HybridAccount{}
	fmt.Println(ha.AvailableFunds())
}
