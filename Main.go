package main

func main() {
	// Initialization of Cards
	var cards []Card
	cards = append(cards, &MasterCard{cardNumber: "4111222233334444", balance: 8000.00, secureCode: 1111})
	cards = append(cards, &VisaCard{cardNumber: "5111222233334444", balance: 6000.00, secureCode: 1112})

	// Initialization of Pizzas
	var pizzaList []*Pizza
	pizzaList = append(pizzaList, &Pizza{name: "- CREATE CUSTOM PIZZA -"})
	pizzaList = append(pizzaList, &Pizza{name: "Philadelphia", size: "Medium", tomato: true, cheese: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "Pepperoni", size: "Small", tomato: true, cheese: true})
	pizzaList = append(pizzaList, &Pizza{name: "Arriva", size: "Medium", tomato: true, sausage: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Hawaiian", size: "Large", pineapple: true, cheese: true, tomato: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Margarita", size: "Large", pepperoni: true, anchovy: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "- END ORDER -"})

	// Creation of Order
	order := &OrdinaryOrder{pizzas: make(map[*Pizza]float32)}
	order.GenerateOrderNum()
	order.SetClient()
	order.PizzaOrdering(pizzaList)
	order.PurchaseProcess(cards)
}
