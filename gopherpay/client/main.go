package main

import (
	"fmt"

	"github.com/jherey/oop-learn/gopherpay/payment"
)

func main() {

	credit := payment.CreditCardAccount(
		"Jeremiah Olufayo",
		"1111-2222-3333-4444",
		2,
		2025,
		1234,
	)

	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	fmt.Printf("Card number: %v\n", credit.CardNumber())
	fmt.Println("Trying to change card number")

	err := credit.SetCardNumber("invalid")
	if err != nil {
		fmt.Printf("That didn't work: %v\n", err)
	}
	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())
}
