package main

// Visitor

type Order interface {
	Accept(Visitor)
}

// Facade
type OrdinaryOrder struct {
	orderNum int
	pizzas   []Pizza //Map should be used
	client   Client
	card     Card // Bridge

}

func (o *OrdinaryOrder) Accept(v Visitor) {
	v.VisitForOrdinaryOrder(o)
}

type Visitor interface {
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
