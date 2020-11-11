package main

import "fmt"

// Bridge DP
type Card interface {
	PayOperation(float32)
}

type VisaCard struct {
	cardNumber string
	balance    float32
	secureCode int
}

func (v *VisaCard) GetCardNumber() string {
	return v.cardNumber
}
func (v *VisaCard) PayOperation(money float32) {
	if money < v.balance {
		v.balance = v.balance - money
		fmt.Println(" - Pay Operation is successfully Completed by Visa Card - ")
	} else {
		fmt.Println(" - Payment Fail, no enough money - ")
	}
}
func (v *VisaCard) GetSecureCode() int {
	return v.secureCode
}

type MasterCard struct {
	cardNumber string
	balance    float32
	secureCode int
}

func (m *MasterCard) GetCardNumber() string {
	return m.cardNumber
}
func (m *MasterCard) PayOperation(money float32) {
	if money < m.balance {
		m.balance = m.balance - money
		fmt.Println(" - Pay Operation is successfully Completed by Master Card - ")
	} else {
		fmt.Println(" - Payment Fail, no enough money - ")
	}
}
func (m *MasterCard) GetSecureCode() int {
	return m.secureCode
}
