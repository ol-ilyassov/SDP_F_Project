package main

import "fmt"

type BasePizza interface {
	GetName() string
	SetName(string)
	GetSize() string
	Accept(PizzaVisitor)
}

// Builder DP
type Pizza struct {
	name                                                            string
	size                                                            string
	tomato, pineapple, anchovy, cheese, pepperoni, lettuce, sausage bool
}

func (p *Pizza) GetName() string { return p.name }

func (p *Pizza) SetName(n string) { p.name = n }

func (p *Pizza) GetSize() string { return p.size }

type PizzaBuilder struct{ pizza *Pizza }

func NewPizzaBuilder() *PizzaBuilder { return &PizzaBuilder{&Pizza{}} }

func (p *PizzaBuilder) Build() *Pizza { return p.pizza }

type SizeBuilder struct{ PizzaBuilder }

func (p *PizzaBuilder) SetSize() *SizeBuilder { return &SizeBuilder{*p} }

func (s *SizeBuilder) Small() *SizeBuilder {
	s.pizza.size = "Small"
	return s
}
func (s *SizeBuilder) Medium() *SizeBuilder {
	s.pizza.size = "Medium"
	return s
}
func (s *SizeBuilder) Large() *SizeBuilder {
	s.pizza.size = "Large"
	return s
}

func (p *PizzaBuilder) AddTomato() *PizzaBuilder {
	p.pizza.tomato = true
	return p
}
func (p *PizzaBuilder) AddPineapple() *PizzaBuilder {
	p.pizza.pineapple = true
	return p
}
func (p *PizzaBuilder) AddAnchovy() *PizzaBuilder {
	p.pizza.anchovy = true
	return p
}
func (p *PizzaBuilder) AddCheese() *PizzaBuilder {
	p.pizza.cheese = true
	return p
}
func (p *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	p.pizza.pepperoni = true
	return p
}
func (p *PizzaBuilder) AddLettuce() *PizzaBuilder {
	p.pizza.lettuce = true
	return p
}
func (p *PizzaBuilder) AddSausage() *PizzaBuilder {
	p.pizza.sausage = true
	return p
}

// Visitor PD
func (p *Pizza) Accept(v PizzaVisitor) {
	v.VisitForPizza(p)
}

type PizzaVisitor interface {
	VisitForPizza(*Pizza)
}

// Visitor Func 1
type GetPrice struct {
	price float32
}

func (gp *GetPrice) ReturnPrice() float32 {
	return gp.price
}
func (gp *GetPrice) VisitForPizza(p *Pizza) {
	gp.price = 0
	switch p.size {
	case "Small":
		gp.price += 100
	case "Medium":
		gp.price += 150
	case "Large":
		gp.price += 200
	}
	if p.tomato {
		gp.price += 150
	}
	if p.pineapple {
		gp.price += 250
	}
	if p.anchovy {
		gp.price += 300
	}
	if p.cheese {
		gp.price += 170
	}
	if p.pepperoni {
		gp.price += 230
	}
	if p.lettuce {
		gp.price += 135
	}
	if p.sausage {
		gp.price += 195
	}
}

// Visitor Func 2
type PrintIngredients struct{}

func (pi *PrintIngredients) VisitForPizza(p *Pizza) {
	message := "Ingredients: "
	if p.tomato {
		message += "Tomato; "
	}
	if p.pineapple {
		message += "Pineapple; "
	}
	if p.anchovy {
		message += "Anchovy; "
	}
	if p.cheese {
		message += "Cheese; "
	}
	if p.pepperoni {
		message += "Pepperoni; "
	}
	if p.lettuce {
		message += "Lettuce; "
	}
	if p.sausage {
		message += "Sausage; "
	}
	fmt.Println(p.GetName() + " with " + message)
}
