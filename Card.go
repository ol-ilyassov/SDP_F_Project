package main

import "fmt"

func (o *OrdinaryOrder) Pay() {
	fmt.Println("- Starting Pay Operation ... -")
	o.card.PayOperation()
}

func (o *OrdinaryOrder) SetCard(card Card) {
	o.card = card
}

// Bridge DP
type Card interface {
	PayOperation()
}

type Visa struct {
	cardNumber int
	balance    float32
}

func (v *Visa) GetCardNumber() int {
	return v.cardNumber
}
func (v *Visa) PayOperation(money float32) {
	if money < v.balance {
		v.balance = v.balance - money
		fmt.Println(" - Pay Operation is successfully Completed by Visa Card - ")
	} else {
		fmt.Println(" - Payment Fail, no enough money - ")
	}
}

type MasterCard struct {
	cardNumber int
	balance    float32
}

func (m *MasterCard) GetCardNumber() int {
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
