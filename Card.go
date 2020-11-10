package main

import "fmt"

type ClientPayer struct {
	card Card
}

func (c *ClientPayer) Pay() {
	fmt.Println("- Starting Pay Operation ... -")
	c.card.PayOperation()
}

func (c *ClientPayer) SetCard(card Card) {
	c.card = card
}

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

func main() {
	visa1 := &Visa{}
	mastercard1 := &MasterCard{}

	client1 := &ClientPayer{}

	client1.SetCard(visa1)
	client1.Pay()
	fmt.Println()

	client1.SetCard(mastercard1)
	client1.Pay()
	fmt.Println()
}
