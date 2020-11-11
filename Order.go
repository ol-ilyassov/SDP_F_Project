package main

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

func (o *OrdinaryOrder) Payment() bool {
	return false
}

/*
func (o *OrdinaryOrder) Accept(v OrderVisitor) {
	v.VisitForOrdinaryOrder(o)
}

type OrderVisitor interface {
	VisitForOrdinaryOrder(*OrdinaryOrder)
}

// Visitor1
type Payment struct{}

func (p *Payment) VisitForOrdinaryOrder(o *OrdinaryOrder) {

}

// Visitor2
type Method2 struct{}

func (m *Method2) VisitForOrdinaryOrder(o *OrdinaryOrder) {

}

//

*/
