package main

import "fmt"

// Visitor
/*
type Order interface {
	Accept(OrderVisitor)
}
*/
// Facade
type OrdinaryOrder struct {
	orderNum int
	pizzas   map[*Pizza]float32 // Pizza Struct + Count
	client   Client
	card     Card // Bridge DP
	isPaid   bool
}

func (o *OrdinaryOrder) SetOrderNum(orderNum int) {
	o.orderNum = orderNum
}

func (o *OrdinaryOrder) GetOrderNum() int {
	return o.orderNum
}

func (o *OrdinaryOrder) SetClient(client Client) {
	o.client = client
}

func (o *OrdinaryOrder) GetClient() *Client {
	return &o.client
}

func (o *OrdinaryOrder) AddPizza(pizza *Pizza, number float32) {
	o.pizzas[pizza] = number
}

func (o *OrdinaryOrder) InitPizzasListForOrder() {
	o.pizzas = make(map[*Pizza]float32)
}

func (o *OrdinaryOrder) Pay(money float32) {
	fmt.Println(" - Starting Pay Operation ... - ")
	o.isPaid = o.card.PayOperation(money)
}

func (o *OrdinaryOrder) SetCard(card Card) {
	o.card = card
}

func (o *OrdinaryOrder) GetCard() *Card {
	return &o.card
}
