package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialization of pizzas
	var pizzaList []*Pizza
	pizzaList = append(pizzaList, &Pizza{name: "Philadelphia", size: "Medium", tomato: true, cheese: true, sausage: true, price: 665})
	pizzaList = append(pizzaList, &Pizza{name: "Pepperoni", size: "Small", tomato: true, cheese: true, price: 420})
	pizzaList = append(pizzaList, &Pizza{name: "Arriva", size: "Medium", tomato: true, sausage: true, lettuce: true, price: 630})
	pizzaList = append(pizzaList, &Pizza{name: "Hawaiian", size: "Large", pineapple: true, cheese: true, tomato: true, lettuce: true, price: 905})
	pizzaList = append(pizzaList, &Pizza{name: "Margarita", size: "Large", pepperoni: true, anchovy: true, sausage: true, price: 925})
	pizzaList = append(pizzaList, &Pizza{name: "CUSTOM"})

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

	fmt.Println(" - Now " + client.GetName() + ", Let's make an Order! - ")
	fmt.Println("> Please, choose a Pizza: ")
	// Modify
	for i, v := range pizzaList {
		fmt.Printf(" [%d] %s\n", i+1, v.GetName())
	}
	// Uncompleted
	fmt.Println("\n - You have chosen the service of buying a Custom Pizza: - ")
	order := NewPizzaBuilder()
	fmt.Printf("> Choose the size for your pizza\n  Small, Medium or Large [S/M/L]: ")
	for {
		fmt.Fscan(os.Stdin, &temp)
		if temp == "S" {
			order.SetSize().Small()
			break
		} else if temp == "M" {
			order.SetSize().Medium()
			break
		} else if temp == "L" {
			order.SetSize().Large()
			break
		} else {
			fmt.Print("> Incorrect size of Pizza, Try again: ")
		}
	}

LoopTomato:
	for {
		fmt.Print("> Would you like to add Tomato? [Y/N]: ")
		switch fmt.Fscan(os.Stdin, &temp); temp {
		case "Y":
			order.AddTomato()
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
			order.AddPineapple()
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
			order.AddAnchovy()
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
			order.AddCheese()
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
			order.AddPepperoni()
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
			order.AddLettuce()
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
			order.AddSausage()
			break LoopSausage
		case "N":
			break LoopSausage
		default:
			fmt.Print(" - Incorrect Input - \n")
		}
	}
	customPizza := order.Build()
	customPizza.ToPrintIngredients()
	fmt.Println("Pizza size: " + customPizza.size)
	fmt.Printf("Total price: $%.2f \n", customPizza.price)

	fmt.Println(" - Thank You for using our Services - ")
	fmt.Println(" - Have a nice Day! - ")
}
