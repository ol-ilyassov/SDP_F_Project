package main

import "fmt"

func main() {
	// Initialization of Cards
	var cards []Card
	cards = append(cards, &MasterCard{cardNumber: "4111222233334444", balance: 800.00, secureCode: 1111})
	cards = append(cards, &VisaCard{cardNumber: "5111222233334444", balance: 600.00, secureCode: 1112})

	// Initialization of Pizzas
	var pizzaList []*Pizza
	pizzaList = append(pizzaList, &Pizza{name: "- CREATE CUSTOM PIZZA -"})
	pizzaList = append(pizzaList, &Pizza{name: "Philadelphia", size: "Medium", tomato: true, cheese: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "Pepperoni", size: "Small", tomato: true, cheese: true})
	pizzaList = append(pizzaList, &Pizza{name: "Arriva", size: "Medium", tomato: true, sausage: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Hawaiian", size: "Large", pineapple: true, cheese: true, tomato: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Margarita", size: "Large", pepperoni: true, anchovy: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "- END ORDER -"})

	// Ordering Process
	fmt.Println(" - Welcome, Dear Client! - ")
	fmt.Println(" - We are glad to see You! - ")
	fmt.Println(" - We want to create a small portrait of You - ")
	stepFour := &PurchaseProcess{}
	stepFour.setCards(cards)

	stepThree := &PizzaOrdering{}
	stepThree.setPizzaList(pizzaList)
	stepThree.setNext(stepFour)

	stepTwo := &CreateClient{}
	stepTwo.setNext(stepThree)

	stepOne := &GenerateOrderNum{}
	stepOne.setNext(stepTwo)

	order := &Order{pizzas: make(map[*Pizza]float32)}
	stepOne.execute(order)

	fmt.Println()
	fmt.Println(" - Thank You for using our Services - ")
	fmt.Println(" - Have a Nice Day! - ")
}
