package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

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
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	order.SetOrderNum(r1.Intn(9999))

	getPrice := &GetPrice{}
	var name, temp string
	var count float32
	var day, month, year, pizzaId, code int

	fmt.Println(" - Welcome, Dear Client! - ")
	fmt.Println(" - We are glad to see You! - ")
	fmt.Println(" - We want to create a small portrait of You - ")
	fmt.Print("> Please, Enter your name: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Println(" - Date of birth may affect on your Future Discount! - ")
	fmt.Print("> Please, Enter your date of birth (Day Month Year): ")
	fmt.Fscan(os.Stdin, &day, &month, &year)
	order.SetClient(NewClientFactory(name, Date{year, month, day}))
	fmt.Println("\n - Now " + order.GetClient().GetName() + ", Let's make an Order! - ")

PizzaOrdering:
	for {
		fmt.Println(" - Pizza's Menu: - ")
		for i, v := range pizzaList {
			fmt.Printf(" [%d] %s\n", i, v.GetName())
		}
		fmt.Println("> Please, choose a PizzaID or Function: ")
		for {
			fmt.Fscan(os.Stdin, &pizzaId)
			if pizzaId < 0 || pizzaId >= len(pizzaList) {
				fmt.Println("> Incorrect PizzaID, Try Again: ")
			} else {
				break
			}
		}
		switch pizzaList[pizzaId].GetName() {
		case "- CREATE CUSTOM PIZZA -":
			pizzaList[0] = order.MakingCustomPizza()
		case "- END ORDER -":
			break PizzaOrdering
		default:
			pizzaList[pizzaId].Accept(&PrintIngredients{})
			pizzaList[pizzaId].Accept(getPrice)
			fmt.Printf("Price: $%.2f\n", getPrice.ReturnPrice())
		PizzaContinueConfirmation:
			for {
				fmt.Println("> Would you like to Continue? [Y/N]: ")
				fmt.Fscan(os.Stdin, &temp)
				switch temp {
				case "Y":
					break PizzaContinueConfirmation
				case "N":
					continue PizzaOrdering
				default:
					fmt.Println(" - Incorrect Input, Try Again - ")
				}
			}
			fmt.Println("> Please, choose a Count of Pizza (Max. 20): ")
			for {
				fmt.Fscan(os.Stdin, &count)
				if count < 1 || count > 20 {
					fmt.Println("> Incorrect count, Change the count: ")
				} else {
					break
				}
			}
			order.AddPizza(pizzaList[pizzaId], count)
		}
		fmt.Println(" |------------------------| ")
		fmt.Printf(" - Order #%d - \n", order.GetOrderNum())
		fmt.Println(" - You Ordered Pizzas: - ")
		for i, v := range order.pizzas {
			i.Accept(getPrice)
			fmt.Printf(" 1) %s - Price: $%.2f - Count: %.0f\n", i.GetName(), getPrice.ReturnPrice(), v)
		}
		fmt.Println(" |------------------------| ")
		fmt.Println()
	}
	fmt.Println(" |------------------------| ")
	fmt.Printf(" - Order #%d - \n", order.GetOrderNum())
	fmt.Println(" - You Ordered Pizzas: - ")
	count = 0
	for i, v := range order.pizzas {
		i.Accept(getPrice)
		fmt.Printf(" 1) %s - Price: $%.2f - Count: %.0f\n", i.GetName(), getPrice.ReturnPrice(), v)
		count++
	}
	if count == 0 {
		fmt.Println("    No Pizza Ordered    ")
	}
	fmt.Println(" |------------------------| ")
	if count != 0 {
		fmt.Println()
		var totalPrice float32
		count = 0
		for i, v := range order.pizzas {
			if v != 0 {
				i.Accept(getPrice)
				totalPrice += getPrice.ReturnPrice() * v
			}
			count++
		}
		fmt.Printf(" - Total Price: $%.2f - \n", totalPrice)
		fmt.Println(order.client.socialStatus)
		count = order.client.MakeDiscount()
		totalPrice = totalPrice - totalPrice*(count/100)
		fmt.Printf(" - Final Payment: $%.2f - \n", totalPrice)
		fmt.Println(" - Please, Provide your Card to make Payment - ")
		// Payment
	CardSetting:
		for {
			fmt.Println("> Enter Your Card Number [or You can enter \"1\" to Exit]: ")
			fmt.Fscan(os.Stdin, &temp)
			if temp == "1" {
				fmt.Println(" - The Payment of Order Refused - ")
				goto End
			}
			for _, v := range cards {
				if temp == v.GetCardNumber() {
					fmt.Print("> Secure Code: ")
					fmt.Fscan(os.Stdin, &code)
					if code == v.GetSecureCode() {
						order.SetCard(v)
						break CardSetting
					} else {
						fmt.Println(" - Incorrect Secure Code - ")
					}
				}
			}
			fmt.Println(" - Incorrect Card data, Try Again - ")
		}
		order.Pay(totalPrice)
		if !order.isPaid {
			goto CardSetting
		}
	}
End:
	fmt.Println()
	fmt.Println(" - Thank You for using our Services - ")
	fmt.Println(" - Have a Nice Day! - ")
}
