package main

import "fmt"

// Bridge DP
type Card interface {
	PayOperation(float32) bool
	GetCardNumber() string
	GetSecureCode() int
}

type VisaCard struct {
	cardNumber string
	balance    float32
	secureCode int
}

func (v *VisaCard) GetCardNumber() string {
	return v.cardNumber
}
func (v *VisaCard) PayOperation(money float32) bool {
	if money < v.balance {
		v.balance = v.balance - money
		fmt.Println(" - Pay Operation is successfully Completed by Visa Card - ")
		return true
	} else {
		fmt.Println(" - Payment Fail, no enough money - ")
	}
	return false
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
func (m *MasterCard) PayOperation(money float32) bool {
	if money < m.balance {
		m.balance = m.balance - money
		fmt.Println(" - Pay Operation is successfully Completed by Master Card - ")
		return true
	} else {
		fmt.Println(" - Payment Fail, no enough money - ")
	}
	return false
}
func (m *MasterCard) GetSecureCode() int {
	return m.secureCode
}
