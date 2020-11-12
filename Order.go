package main

import (
	"fmt"
	"os"
)

// Facade
type OrdinaryOrder struct {
	orderNum int
	pizzas   map[*Pizza]float32 // PizzaStruct + Count
	client   *Client
	card     Card // Bridge DP
	isPaid   bool
}

func (o *OrdinaryOrder) SetOrderNum(orderNum int) {
	o.orderNum = orderNum
}

func (o *OrdinaryOrder) GetOrderNum() int {
	return o.orderNum
}

func (o *OrdinaryOrder) SetClient(client *Client) {
	o.client = client
}

func (o *OrdinaryOrder) GetClient() *Client {
	return o.client
}

func (o *OrdinaryOrder) AddPizza(pizza *Pizza, number float32) {
	o.pizzas[pizza] = number
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

// Complex Functions

func (o *OrdinaryOrder) MakingCustomPizza() *Pizza {
	var temp string
	fmt.Println(" - You have chosen the service of buying a Custom Pizza: - ")
	fmt.Println(" - You should choose 2 or more ingredients - ")
	pizzaBuilder := NewPizzaBuilder()
	fmt.Printf("> Choose the size for your pizza\n  Small, Medium, or Large [S/M/L]: ")
	for {
		fmt.Fscan(os.Stdin, &temp)
		if temp == "S" {
			pizzaBuilder.SetSize().Small()
			break
		} else if temp == "M" {
			pizzaBuilder.SetSize().Medium()
			break
		} else if temp == "L" {
			pizzaBuilder.SetSize().Large()
			break
		} else {
			fmt.Print("> Incorrect size of Pizza, Try again: ")
		}
	}
	var tempBuilder *PizzaBuilder
	var count int
	for {
		tempBuilder = pizzaBuilder
		count = 0

	LoopTomato:
		for {
			fmt.Print("> Would you like to add Tomato? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddTomato()
				count += 1
				break LoopTomato
			case "N":
				break LoopTomato
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopPineapple:
		for {
			fmt.Print("> Would you like to add Pineapple? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddPineapple()
				count += 1
				break LoopPineapple
			case "N":
				break LoopPineapple
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopAnchovy:
		for {
			fmt.Print("> Would you like to add Anchovy? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddAnchovy()
				count += 1
				break LoopAnchovy
			case "N":
				break LoopAnchovy
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopCheese:
		for {
			fmt.Print("> Would you like to add Cheese? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddCheese()
				count += 1
				break LoopCheese
			case "N":
				break LoopCheese
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopPepperoni:
		for {
			fmt.Print("> Would you like to add Pepperoni? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddPepperoni()
				count += 1
				break LoopPepperoni
			case "N":
				break LoopPepperoni
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopLettuce:
		for {
			fmt.Print("> Would you like to add Lettuce? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddLettuce()
				count += 1
				break LoopLettuce
			case "N":
				break LoopLettuce
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopSausage:
		for {
			fmt.Print("> Would you like to add Sausage? [Y/N]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "Y":
				tempBuilder.AddSausage()
				count += 1
				break LoopSausage
			case "N":
				break LoopSausage
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
		if count >= 2 {
			break
		}
		fmt.Println(" - You have chosen less than 2 ingredients for Pizza - ")
		fmt.Println(" - Retry choosing the ingredients for Pizza - ")
	}
	pizzaBuilder = tempBuilder
	customPizza := pizzaBuilder.Build()
	customPizza.SetName("Custom Pizza")
	customPizza.Accept(&PrintIngredients{})
	fmt.Println("Custom Pizza size: " + customPizza.GetSize())
	getPrice := &GetPrice{}
	customPizza.Accept(getPrice)
	fmt.Printf("Custom Pizza price: $%.2f \n", getPrice.price)
	return customPizza
}
