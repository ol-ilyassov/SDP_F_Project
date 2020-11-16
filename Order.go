package main

import (
	"fmt"
)

type Order struct {
	orderNum int
	pizzas   map[*Pizza]float32 // PizzaStruct + Count
	client   *Client
	card     Card // Bridge DP
	isPaid   bool

	orderNumGeneration bool
	clientCreation     bool
	pizzaOrdering      bool
	purchaseProcess    bool
}

func (o *Order) GetOrderNum() int {
	return o.orderNum
}
func (o *Order) SetOrderNum(num int) {
	o.orderNum = num
}
func (o *Order) GetClient() *Client {
	return o.client
}
func (o *Order) GetPizzasList() map[*Pizza]float32 {
	return o.pizzas
}
func (o *Order) AddPizza(pizza *Pizza, number float32) {
	o.pizzas[pizza] = number
}
func (o *Order) SetCard(card Card) {
	o.card = card
}
func (o *Order) GetCard() *Card {
	return &o.card
}
func (o *Order) Pay(money float32) {
	fmt.Println(" - Starting Pay Operation ... - ")
	o.isPaid = o.card.PayOperation(money)
}
