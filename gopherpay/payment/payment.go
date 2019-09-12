package payment

import (
	"errors"
	"regexp"
	"time"
)

// CreditCard data structure
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreditCardAccount constructor function to create a CreditCard
func CreditCardAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// OwnerName retuns an account's owner's name
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName sets an account's owner's name
func (c *CreditCard) SetOwnerName(name string) error {
	if len(name) == 0 {
		return errors.New("Invalid owner name provided")
	}
	c.ownerName = name
	return nil
}

// CardNumber retuns an account's card number
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

// SetCardNumber sets the card number of the account
func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}
	c.cardNumber = value
	return nil
}

// ExpirationDate returns an account's expiration date
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate sets the expiration data for an account
func (c CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode retuns an account's security code
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode sets the secuiry code for an account
func (c *CreditCard) SetSecurityCode(code int) {
	if code < 100 || code > 999 {
	}
}

// AvailableCredit retuns an account's available credit
func (c CreditCard) AvailableCredit() float32 {
	return 5000. // this can come from a webservice, client doesn't know or care
}
