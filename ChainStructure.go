package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type OrderingStep interface {
	execute(*Order)
	setNext(OrderingStep)
}

type GenerateOrderNum struct {
	next OrderingStep
}

func (g *GenerateOrderNum) execute(o *Order) {
	if o.orderNumGeneration {
		fmt.Println(" - \"Order Number Generation\" Step is already Completed - ")
		g.next.execute(o)
		return
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	o.SetOrderNum(r1.Intn(9999))
	o.orderNumGeneration = true
	g.next.execute(o)
}
func (g *GenerateOrderNum) setNext(next OrderingStep) {
	g.next = next
}

type CreateClient struct {
	next OrderingStep
}

func (c *CreateClient) execute(o *Order) {
	if o.clientCreation {
		fmt.Println(" - \"Client Creation Step\" is already Completed - ")
		c.next.execute(o)
		return
	}
	var name string
	var day, month, year int
	fmt.Print("> Please, Enter your name: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Println(" - Date of birth may affect on your Future Discount! - ")
	for {
		fmt.Print("> Please, Enter your Day of birth (Valid Input Day: 1-31): ")
		fmt.Fscan(os.Stdin, &day)
		if day < 1 || day > 31 {
			fmt.Println(" - Incorrect Input about Day - ")
			continue
		}
		break
	}
	for {
		fmt.Print("> Please, Enter your Month of birth (Valid Input: 1-12): ")
		fmt.Fscan(os.Stdin, &month)
		if month < 1 || month > 12 {
			fmt.Println(" - Incorrect Input about Month - ")
			continue
		}
		break
	}
	for {
		fmt.Print("> Please, Enter your Year of birth (Valid Input Year: from 1900 to 2015): ")
		fmt.Fscan(os.Stdin, &year)
		if year < 1900 || year > 2015 {
			fmt.Println(" - Incorrect Input about Year - ")
			continue
		}
		break
	}
	o.SetClient(NewClientFactory(name, year, month, day))
	o.clientCreation = true
	c.next.execute(o)
}
func (c *CreateClient) setNext(next OrderingStep) {
	c.next = next
}

type PizzaOrdering struct {
	next      OrderingStep
	pizzaList []*Pizza
}

func (p *PizzaOrdering) execute(o *Order) {
	if o.pizzaOrdering {
		fmt.Println(" - \"Pizza Ordering Step\" is already Completed - ")
		p.next.execute(o)
		return
	}
	getPrice := &GetPrice{}
	var count float32
	var pizzaId int
	var temp string

	fmt.Println("\n - Now " + o.GetClient().GetName() + ", Let's make an Order! - ")
PizzaOrdering:
	for {
		fmt.Println(" - Pizza's Menu: - ")
		for i, v := range p.pizzaList {
			fmt.Printf(" [%d] %s\n", i+1, v.GetName())
		}
		fmt.Println("> Please, choose a PizzaID or Function: ")
		for {
			fmt.Fscan(os.Stdin, &pizzaId)
			if pizzaId < 1 || pizzaId >= len(p.pizzaList)+1 {
				fmt.Println("> Incorrect PizzaID or Function, Try Again: ")
			} else {
				break
			}
		}
		pizzaId--
		switch p.pizzaList[pizzaId].GetName() {
		case "- CREATE CUSTOM PIZZA -":
			p.pizzaList[0] = p.MakingCustomPizza()
		case "- END ORDER -":
			break PizzaOrdering
		default:
			p.pizzaList[pizzaId].Accept(&PrintIngredients{})
			p.pizzaList[pizzaId].Accept(getPrice)
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
			o.AddPizza(p.pizzaList[pizzaId], count)
		}
		fmt.Println(" |------------------------| ")
		fmt.Printf(" - Order #%d - \n", o.GetOrderNum())
		fmt.Println(" - You Ordered Pizzas: - ")
		for i, v := range o.GetPizzasList() {
			i.Accept(getPrice)
			fmt.Printf(" 1) %s - Price: $%.2f - Count: %.0f\n", i.GetName(), getPrice.ReturnPrice(), v)
		}
		fmt.Println(" |------------------------| ")
		fmt.Println()
	}
	o.pizzaOrdering = true
	p.next.execute(o)
}
func (p *PizzaOrdering) setNext(next OrderingStep) {
	p.next = next
}
func (p *PizzaOrdering) setPizzaList(list []*Pizza) {
	p.pizzaList = list
}
func (p *PizzaOrdering) MakingCustomPizza() *Pizza {
	var temp string
	fmt.Println(" - You have chosen the service of buying a Custom Pizza: - ")
	fmt.Println(" - You should choose 2 or more ingredients - ")
	pizzaBuilder := NewPizzaBuilder()
	fmt.Printf("> Choose the size for your pizza\n  Small, Medium, or Large [s/m/l]: ")
	for {
		fmt.Fscan(os.Stdin, &temp)
		if temp == "s" {
			pizzaBuilder.SetSize().Small()
			break
		} else if temp == "m" {
			pizzaBuilder.SetSize().Medium()
			break
		} else if temp == "l" {
			pizzaBuilder.SetSize().Large()
			break
		} else {
			fmt.Print("> Incorrect size of Pizza, Try again: ")
		}
	}
	var intermediateBuilder *PizzaBuilder
	var count int
	for {
		intermediateBuilder = pizzaBuilder
		count = 0

	LoopTomato:
		for {
			fmt.Print("> Would you like to add Tomato? [y/n]: ")
			switch fmt.Fscan(os.Stdin, &temp); temp {
			case "y":
				intermediateBuilder.AddTomato()
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
				intermediateBuilder.AddPineapple()
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
				intermediateBuilder.AddAnchovy()
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
				intermediateBuilder.AddCheese()
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
				intermediateBuilder.AddPepperoni()
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
				intermediateBuilder.AddLettuce()
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
				intermediateBuilder.AddSausage()
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
	fmt.Println()
	pizzaBuilder = intermediateBuilder
	customPizza := pizzaBuilder.Build()
	customPizza.SetName("Custom Pizza")
	customPizza.Accept(&PrintIngredients{})
	fmt.Println("Custom Pizza size: " + customPizza.GetSize())
	getPrice := &GetPrice{}
	customPizza.Accept(getPrice)
	fmt.Printf("Price for Custom Pizza: $%.2f \n", getPrice.price)
	fmt.Printf(" - Now, You can order Custom Pizza - \n")
	return customPizza
}

type PurchaseProcess struct {
	next  OrderingStep
	cards []Card
}

func (p *PurchaseProcess) execute(o *Order) {
	if o.purchaseProcess {
		fmt.Println(" - \"Purchase Process Step\" is already Completed - ")
		p.next.execute(o)
		return
	}
	getPrice := &GetPrice{}
	var temp string
	var count float32
	var code int

	fmt.Println(" |------------------------| ")
	fmt.Printf(" - Order #%d - \n", o.GetOrderNum())
	fmt.Println(" - You Ordered Pizzas: - ")
	count = 0
	for i, v := range o.GetPizzasList() {
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
		for i, v := range o.GetPizzasList() {
			if v != 0 {
				i.Accept(getPrice)
				totalPrice += getPrice.ReturnPrice() * v
			}
			count++
		}
		fmt.Printf(" - Total Price: $%.2f - \n", totalPrice)
		fmt.Println(o.GetClient().GetStatus())
		count = o.GetClient().MakeDiscount()
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
			for _, v := range p.cards {
				if temp == v.GetCardNumber() {
					fmt.Print("> Secure Code: ")
					fmt.Fscan(os.Stdin, &code)
					if v.ValidateSecureCode(code) {
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
func (p *PurchaseProcess) setNext(next OrderingStep) {
	p.next = next
}
func (p *PurchaseProcess) setCards(list []Card) {
	p.cards = list
}
