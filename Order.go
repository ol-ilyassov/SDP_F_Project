package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Facade
type OrdinaryOrder struct {
	orderNum int
	pizzas   map[*Pizza]float32 // PizzaStruct + Count
	client   *Client
	card     Card // Bridge DP
	isPaid   bool
}

func (o *OrdinaryOrder) GetOrderNum() int {
	return o.orderNum
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

// #1 Client Creation
func (o *OrdinaryOrder) CreateClient() {
	var name string
	var day, month, year int

	fmt.Print("> Please, Enter your name: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Println(" - Date of birth may affect on your Future Discount! - ")
	for {
		fmt.Print("> Please, Enter your date of birth (dd mm yyyy): ")
		fmt.Fscan(os.Stdin, &day, &month, &year)
		if day < 1 || day > 31 {
			fmt.Println(" - Incorrect Input about Day (Valid Input Day: 1-31) - ")
			continue
		}
		if month < 1 || month > 12 {
			fmt.Println(" - Incorrect Input about Month (Valid Input Day: 1-12) - ")
			continue
		}
		if year < 1950 || year > 2015 {
			fmt.Println(" - Incorrect Input about Year (Valid Input Year: from 1950 to 2015) - ")
			continue
		}
		break
	}

	o.client = NewClientFactory(name, year, month, day)
}

// #3 Pizza Ordering
func (o *OrdinaryOrder) PizzaOrdering(pizzaList []*Pizza) {
	getPrice := &GetPrice{}
	var count float32
	var pizzaId int
	var temp string

	fmt.Println("\n - Now " + o.GetClient().GetName() + ", Let's make an Order! - ")
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
			pizzaList[0] = o.MakingCustomPizza()
		case "- END ORDER -":
			break PizzaOrdering
		default:
			pizzaList[pizzaId].Accept(&PrintIngredients{})
			pizzaList[pizzaId].Accept(getPrice)
			fmt.Printf("Price: $%.2f\n", getPrice.ReturnPrice())
		PizzaContinueConfirmation:
			for {
				fmt.Println("> Would you like to Continue? [y/n]: ")
				fmt.Fscan(os.Stdin, &temp)
				switch temp {
				case "y":
					break PizzaContinueConfirmation
				case "n":
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
			o.AddPizza(pizzaList[pizzaId], count)
		}
		fmt.Println(" |------------------------| ")
		fmt.Printf(" - Order #%d - \n", o.GetOrderNum())
		fmt.Println(" - You Ordered Pizzas: - ")
		for i, v := range o.pizzas {
			i.Accept(getPrice)
			fmt.Printf(" 1) %s - Price: $%.2f - Count: %.0f\n", i.GetName(), getPrice.ReturnPrice(), v)
		}
		fmt.Println(" |------------------------| ")
		fmt.Println()
	}
}

// #2 OrderID
func (o *OrdinaryOrder) GenerateOrderNum() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	o.orderNum = r1.Intn(9999)
}

// #4 Custom Pizza Creation
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
			fmt.Print("> Would you like to add Tomato? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddTomato()
				count += 1
				break LoopTomato
			case "n":
				break LoopTomato
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopPineapple:
		for {
			fmt.Print("> Would you like to add Pineapple? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddPineapple()
				count += 1
				break LoopPineapple
			case "n":
				break LoopPineapple
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopAnchovy:
		for {
			fmt.Print("> Would you like to add Anchovy? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddAnchovy()
				count += 1
				break LoopAnchovy
			case "n":
				break LoopAnchovy
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopCheese:
		for {
			fmt.Print("> Would you like to add Cheese? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddCheese()
				count += 1
				break LoopCheese
			case "n":
				break LoopCheese
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopPepperoni:
		for {
			fmt.Print("> Would you like to add Pepperoni? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddPepperoni()
				count += 1
				break LoopPepperoni
			case "n":
				break LoopPepperoni
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopLettuce:
		for {
			fmt.Print("> Would you like to add Lettuce? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddLettuce()
				count += 1
				break LoopLettuce
			case "n":
				break LoopLettuce
			default:
				fmt.Print(" - Incorrect Input - \n")
			}
		}
	LoopSausage:
		for {
			fmt.Print("> Would you like to add Sausage? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				tempBuilder.AddSausage()
				count += 1
				break LoopSausage
			case "n":
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

// #5 Purchase Process
func (o *OrdinaryOrder) PurchaseProcess(cards []Card) {
	getPrice := &GetPrice{}
	var temp string
	var count float32
	var code int

	fmt.Println(" |------------------------| ")
	fmt.Printf(" - Order #%d - \n", o.GetOrderNum())
	fmt.Println(" - You Ordered Pizzas: - ")
	count = 0
	for i, v := range o.pizzas {
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
		for i, v := range o.pizzas {
			if v != 0 {
				i.Accept(getPrice)
				totalPrice += getPrice.ReturnPrice() * v
			}
			count++
		}
		fmt.Printf(" - Total Price: $%.2f - \n", totalPrice)
		fmt.Println(o.client.socialStatus)
		count = o.client.MakeDiscount()
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
						o.SetCard(v)
						break CardSetting
					} else {
						fmt.Println(" - Incorrect Secure Code - ")
					}
				}
			}
			fmt.Println(" - Incorrect Card data, Try Again - ")
		}
		o.Pay(totalPrice)
		if !o.isPaid {
			goto CardSetting
		}
	}
End:
}
