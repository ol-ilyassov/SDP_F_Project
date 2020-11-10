package main

// Visitor

// Facade
type OrdinaryOrder struct {
	orderNum int
	pizzas   []Pizza //Map should be used
	client   Client
	card     Card // Bridge
}
