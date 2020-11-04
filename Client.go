package main

import "fmt"

type client struct {
	name         string
	age          int
	socialStatus status
}

func (c *client) pay() {
	c.socialStatus.discount()
}

func NewClient(name string, age int) *client {
	if age > 50 {
		return &client{name, age, PensionerStatus{}}
	}
	return &client{name, age, DefaultStatus{}}
}

func (c *client) setStatus(s status) {
	c.socialStatus = s
}

type status interface {
	discount()
}

type DefaultStatus struct{}

func (d DefaultStatus) discount() {
	fmt.Println("Your discount is 0%")
}

type StudentStatus struct{}

func (s StudentStatus) discount() {
	fmt.Println("Your discount is 20%")
}

type PensionerStatus struct{}

func (p PensionerStatus) discount() {
	fmt.Println("Your discount is 30%")
}

func main() {
	newclient := NewClient("Olzhas", 20)
	newclient.pay()
	newclient.setStatus(StudentStatus{})
	newclient.pay()
}
