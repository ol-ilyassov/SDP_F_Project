package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialization of pizzas
	var pizzaList []*Pizza
	pizzaList = append(pizzaList, &Pizza{name: "CUSTOM PIZZA"})
	pizzaList = append(pizzaList, &Pizza{name: "Philadelphia", size: "Medium", tomato: true, cheese: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "Pepperoni", size: "Small", tomato: true, cheese: true})
	pizzaList = append(pizzaList, &Pizza{name: "Arriva", size: "Medium", tomato: true, sausage: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Hawaiian", size: "Large", pineapple: true, cheese: true, tomato: true, lettuce: true})
	pizzaList = append(pizzaList, &Pizza{name: "Margarita", size: "Large", pepperoni: true, anchovy: true, sausage: true})
	pizzaList = append(pizzaList, &Pizza{name: "END ORDER"})

	var name, temp string
	var day, month, year int

	fmt.Println(" - Welcome, Dear Client! - ")
	fmt.Println(" - We are glad to see You! - ")
	fmt.Println(" - We want to create a small portrait of You - ")
	fmt.Print("> Please, Enter your name: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Println(" - Date of birth may affect on your Future Discount! - ")
	fmt.Print("> Please, Enter your date of birth (Day Month Year): ")
	fmt.Fscan(os.Stdin, &day, &month, &year)
	client := NewClientFactory(name, Date{year, month, day})

	fmt.Println("\n - Now " + client.GetName() + ", Let's make an Order! - ")
	fmt.Println("> Please, choose a Pizza: ")
	// Modify
	for i, v := range pizzaList {
		fmt.Printf(" [%d] %s\n", i+1, v.GetName())
	}
	// Uncompleted

	fmt.Println("\n - You have chosen the service of buying a Custom Pizza: - ")
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
	}
	pizzaBuilder = tempBuilder
	customPizza := pizzaBuilder.Build()
	customPizza.SetName("Custom Pizza")
	customPizza.Accept(&PrintIngredients{})
	fmt.Println("Custom Pizza size: " + customPizza.GetSize())
	getPrice := &GetPrice{}
	customPizza.Accept(getPrice)
	fmt.Printf("Custom Pizza price: $%.2f \n", getPrice.price)

	//
	// Payment
	//

	fmt.Println(" - Thank You for using our Services - ")
	fmt.Println(" - Have a nice Day! - ")
}
