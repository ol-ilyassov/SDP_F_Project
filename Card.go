package main

import "fmt"

// Bridge DP
type Card interface {
	GetCardNumber() string
	ValidateSecureCode(int) bool
	PayOperation(float32) bool
}

type VisaCard struct {
	cardNumber string
	secureCode int
	balance    float32
}

func (v *VisaCard) GetCardNumber() string {
	return v.cardNumber
}
func (v *VisaCard) ValidateSecureCode(secureCode int) bool {
	if v.secureCode == secureCode {
		return true
	}
	return false
}
func (v *VisaCard) PayOperation(money float32) bool {
	if money < v.balance {
		v.balance = v.balance - money
		fmt.Println(" - Pay Operation is Successfully Completed by Visa Card - ")
		return true
	} else {
		fmt.Println(" - Payment Fail, No enough Money on Balance - ")
	}
	return false
}

type MasterCard struct {
	cardNumber string
	secureCode int
	balance    float32
}

func (m *MasterCard) GetCardNumber() string {
	return m.cardNumber
}
func (m *MasterCard) ValidateSecureCode(secureCode int) bool {
	if m.secureCode == secureCode {
		return true
	}
	return false
}
func (m *MasterCard) PayOperation(money float32) bool {
	if money < m.balance {
		m.balance = m.balance - money
		fmt.Println(" - Pay Operation is Successfully Completed by Master Card - ")
		return true
	} else {
		fmt.Println(" - Payment Fail, No enough Money on Balance - ")
	}
	return false
}
