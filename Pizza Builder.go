// Group: SE1902
// Student 1: Ilyassov Olzhas
// Student 2: Sovetkazhiyev Alibek
// Student 3: Bakhytbekov Yersultan
// Task 1: Group Project
package main

import "fmt"

type Pizza struct {
	name                                                            string
	size                                                            string
	tomato, pineapple, anchovy, cheese, pepperoni, lettuce, sausage bool
	price                                                           float32
}

func (p *Pizza) GetName() string {
	return p.name
}

type PizzaBuilder struct{ pizza *Pizza }

func NewPizzaBuilder() *PizzaBuilder { return &PizzaBuilder{&Pizza{}} }

func (p *PizzaBuilder) Build() *Pizza { return p.pizza }

type SizeBuilder struct{ PizzaBuilder }

func (p *PizzaBuilder) SetSize() *SizeBuilder { return &SizeBuilder{*p} }

func (s *SizeBuilder) Small() *SizeBuilder {
	s.pizza.size = "Small"
	s.pizza.price += 100
	return s
}

func (s *SizeBuilder) Medium() *SizeBuilder {
	s.pizza.size = "Medium"
	s.pizza.price += 150
	return s
}

func (s *SizeBuilder) Large() *SizeBuilder {
	s.pizza.size = "Large"
	s.pizza.price += 200
	return s
}

func (p *PizzaBuilder) AddTomato() *PizzaBuilder {
	p.pizza.tomato = true
	p.pizza.price += 150
	return p
}

func (p *PizzaBuilder) AddPineapple() *PizzaBuilder {
	p.pizza.pineapple = true
	p.pizza.price += 250
	return p
}

func (p *PizzaBuilder) AddAnchovy() *PizzaBuilder {
	p.pizza.anchovy = true
	p.pizza.price += 300
	return p
}

func (p *PizzaBuilder) AddCheese() *PizzaBuilder {
	p.pizza.cheese = true
	p.pizza.price += 170
	return p
}

func (p *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	p.pizza.pepperoni = true
	p.pizza.price += 230
	return p
}

func (p *PizzaBuilder) AddLettuce() *PizzaBuilder {
	p.pizza.lettuce = true
	p.pizza.price += 135
	return p
}

func (p *PizzaBuilder) AddSausage() *PizzaBuilder {
	p.pizza.sausage = true
	p.pizza.price += 195
	return p
}

func (p *Pizza) ToPrintIngredients() {
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
	fmt.Println("Pizza" + p.GetName() + "with " + message)
}

func main() {

	fmt.Println("- Welcome to \"BigO\" Pizzeria")
	fmt.Println("- Your Custom-Pizza: ")

	order := NewPizzaBuilder()
	order.SetSize().Medium().AddAnchovy().AddTomato().AddSausage()
	customPizza := order.Build()

	customPizza.ToPrintIngredients()
	fmt.Println("Pizza size: " + customPizza.size)
	fmt.Printf("Total price: $%.2f \n", customPizza.price)
}
