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
	cardNumber     int
	balance        float32
	encryptionCode int
}

func (v *Visa) PayOperation() {
	fmt.Println("Pay Operation is successfully Completed by Visa Card")
}

type MasterCard struct {
	cardNumber     int
	balance        float32
	encryptionCode int
}

func (p *MasterCard) PayOperation() {
	fmt.Println("Pay Operation is successfully Completed by Master Card")
}
