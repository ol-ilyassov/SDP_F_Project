// Group: SE1902
// Student 1: Ilyassov Olzhas
// Student 2: Sovetkazhiyev Alibek
// Student 3: Bakhytbekov Yersultan
// Task 1: Group Project

package main

/*
import "fmt"

// 0) Concrete Struct
type Pizza struct {
	Name                    string
	Ingredients             []string
	DoughPrepared           bool
	TomatoAndCheesePrepared bool
	CookingDone             bool
}

func (p *Pizza) String() string {
	responseString := "Pizza: " + p.Name
	if p.CookingDone {
		responseString += " is Cooked"
	}
	if p.Ingredients == nil {
		return responseString
	} else {
		responseString += ", with Ingredients:\n"
		for i, v := range p.Ingredients {
			responseString += fmt.Sprintf("  %d) %s. \n", i+1, v)
		}
	}
	return responseString
}
func NewPizza(name string) *Pizza {
	return &Pizza{Name: name}
}

// 1) Handler Interface
type Cooking interface {
	Add(m Cooking)
	Apply()
}

// 2) Base Handler
type CookingStep struct {
	pizza *Pizza
	next  Cooking
}

func (c *CookingStep) Add(m Cooking) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}
func (c *CookingStep) Apply() {
	if c.next != nil {
		c.next.Apply()
	}
}
func newCookingStep(p *Pizza) *CookingStep {
	return &CookingStep{pizza: p}
}

// 3) Concrete Handler
type AddIngredient struct {
	CookingStep
	ingredient string
}

func (a *AddIngredient) Apply() {
	fmt.Println("- Adding:", a.ingredient, "to", a.pizza.Name)
	a.pizza.Ingredients = append(a.pizza.Ingredients, a.ingredient)
	a.CookingStep.Apply()
}
func NewAddIngredient(p *Pizza, ingredient string) *AddIngredient {
	return &AddIngredient{CookingStep{pizza: p}, ingredient}
}

// 4) Concrete Handler
type PrepareDough struct {
	CookingStep
}

func (i *PrepareDough) Apply() {
	if i.pizza.DoughPrepared {
		fmt.Println("- Dough is already prepared for", i.pizza.Name)
		i.CookingStep.Apply()
		return
	}
	fmt.Println("- Preparing Dough for", i.pizza.Name)
	i.pizza.DoughPrepared = true
	i.CookingStep.Apply()
}
func NewPrepareDough(p *Pizza) *PrepareDough {
	return &PrepareDough{CookingStep{pizza: p}}
}

// 5) Concrete Handler
type PrepareTomatoSauceAndCheese struct {
	CookingStep
}

func (i *PrepareTomatoSauceAndCheese) Apply() {
	if i.pizza.TomatoAndCheesePrepared {
		fmt.Println("- Tomato Sauce and Cheese is already prepared for", i.pizza.Name)
		i.CookingStep.Apply()
		return
	}
	fmt.Println("- Preparing Tomato Sauce and Cheese for", i.pizza.Name)
	i.pizza.TomatoAndCheesePrepared = true
	i.CookingStep.Apply()
}
func NewPrepareTomatoSauceAndCheese(p *Pizza) *PrepareTomatoSauceAndCheese {
	return &PrepareTomatoSauceAndCheese{CookingStep{pizza: p}}
}

// 6) Concrete Handler
type CookingInTheOven struct {
	CookingStep
}

func (i *CookingInTheOven) Apply() {
	if i.pizza.CookingDone {
		fmt.Println("- Cooking is already done for", i.pizza.Name)
		return
	}
	fmt.Println("- Finally, Cooking of", i.pizza.Name, "in the Oven")
	i.pizza.CookingDone = true
}
func NewCookingInTheOven(p *Pizza) *CookingInTheOven {
	return &CookingInTheOven{CookingStep{pizza: p}}
}

func main() {
	myPizza := NewPizza("Margherita")
	fmt.Println(myPizza)

	root := newCookingStep(myPizza)
	root.Add(NewPrepareDough(myPizza))
	root.Add(NewPrepareTomatoSauceAndCheese(myPizza))
	root.Add(NewAddIngredient(myPizza, "Basil"))
	root.Add(NewAddIngredient(myPizza, "Mushrooms"))
	root.Add(NewAddIngredient(myPizza, "Olives"))
	root.Add(NewPrepareDough(myPizza))
	root.Add(NewCookingInTheOven(myPizza))

	//No effect to Pizza after Cooking In The Oven process.
	root.Add(NewAddIngredient(myPizza, "Chicken Meat"))

	root.Apply()
	fmt.Println(myPizza)
}
*/
