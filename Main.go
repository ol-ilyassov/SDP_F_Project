package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func MakingCustomPizza() *Pizza {
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

func main() {

	//card1 := &MasterCard{cardNumber: 4515345294534234, balance: 8000.00}

	// Initialization of pizzas
	var pizzaList []*Pizza
	pizzaList = append(pizzaList, &Pizza{name: "- CREATE CUSTOM PIZZA -"})
	pizzaList = append(pizzaList, &Pizza{name: "Philadelphia", size: "Medium", tomato: true, cheese: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "Pepperoni", size: "Small", tomato: true, cheese: true})
	pizzaList = append(pizzaList, &Pizza{name: "Arriva", size: "Medium", tomato: true, sausage: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Hawaiian", size: "Large", pineapple: true, cheese: true, tomato: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Margarita", size: "Large", pepperoni: true, anchovy: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "- END ORDER -"})
	// End

	order := &OrdinaryOrder{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	order.SetOrderNum(r1.Intn(9999))

	var name, temp string
	var count float32
	var day, month, year, pizzaId int

	fmt.Println(" - Welcome, Dear Client! - ")
	fmt.Println(" - We are glad to see You! - ")
	fmt.Println(" - We want to create a small portrait of You - ")
	fmt.Print("> Please, Enter your name: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Println(" - Date of birth may affect on your Future Discount! - ")
	fmt.Print("> Please, Enter your date of birth (Day Month Year): ")
	fmt.Fscan(os.Stdin, &day, &month, &year)
	client := NewClientFactory(name, Date{year, month, day})
	order.SetClient(*client)

	fmt.Println("\n - Now " + client.GetName() + ", Let's make an Order! - ")
	fmt.Println(" - Pizza's Menu: - ")

PizzaOrdering:
	for {
		for i, v := range pizzaList {
			fmt.Printf(" [%d] %s\n", i, v.GetName())
		}
		fmt.Println("> Please, choose a PizzaID or Function: ")
		/*
			x := someFunction() // Some value of an unknown type is stored in x now
			switch x := x.(type) {
			    case bool:
			        fmt.Printf("boolean %t\n", x)             // x has type bool
			    case int:
			        fmt.Printf("integer %d\n", x)             // x has type int
			    case string:
			        fmt.Printf("pointer to boolean %s\n", x) // x has type string
			    default:
			        fmt.Printf("unexpected type %T\n", x)     // %T prints whatever type x is
			}
		*/
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
			pizzaList[0] = MakingCustomPizza()
		case "- END ORDER -":
			break PizzaOrdering
		default:
			pizzaList[pizzaId].Accept(&PrintIngredients{})
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
			order.pizzas[pizzaList[pizzaId]] = count
		}
		fmt.Println(" |------------------------| ")
		fmt.Println(" - You Ordered Pizzas: - ")
		for i, v := range order.pizzas {
			fmt.Printf("\n 1) %s - %f", i.GetName(), v)
		}
	}

	//

	// Payment

	//

	fmt.Println(" - Thank You for using our Services - ")
	fmt.Println(" - Have a nice Day! - ")

	/*
		for i, v := range pizzaList {
			fmt.Printf(" [%d] %s\n", i, v.GetName())
		}
		fmt.Fscan(os.Stdin, &pizzaId)
	*/

	//
	/*
		order.pizzas = make(map[*Pizza]float32)
		order.client = Client{}
		order.pizzas[pizzaList[1]] = 2
		order.pizzas[pizzaList[2]] = 4

		getPrice := &GetPrice{}
		var totalPrice float32
		count = 0
		for i, v := range order.pizzas {
			if v != 0 {
				fmt.Println(order.pizzas[i])
				i.Accept(getPrice)
				totalPrice += getPrice.price * v
				fmt.Println()
			}
			count ++
		}
		fmt.Printf("Custom Pizza price: $%.2f \n", totalPrice)

		pizzaList[0] = &Pizza{name: "Meow", size: "Large",tomato: true, cheese: true}

		fmt.Println(pizzaList[0])
		fmt.Println(pizzaList[1])
	*/
}
