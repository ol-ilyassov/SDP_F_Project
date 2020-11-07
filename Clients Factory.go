package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}
func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}
func main() {
	developerFactory := NewEmployeeFactory("developer", 80000)
	managerFactory := NewEmployeeFactory("manager", 100000)
	bossFactory := NewEmployeeFactory2("CEO", 1500000)
	developer := developerFactory("Adam")
	manager := managerFactory("Bob")
	bossFactory.AnnualIncome = 8000000
	boss := bossFactory.Create("Sam")
	boss.AnnualIncome = 5
	fmt.Println(developer, manager, boss)
}
