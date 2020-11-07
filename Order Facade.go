package main

// Iterator ?
// Visitor ?

type Order interface {
}

type OrdinaryOrder struct {
	orderNum int
	pizzas   []Pizza //Map should be used
	date     string
	cardNum  int
}
