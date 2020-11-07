// Group: SE1902
// Student 1: Ilyassov Olzhas
// Student 2: Sovetkazhiyev Alibek
// Student 3: Bakhytbekov Yersultan
// Task 1: Group Project
package main

import "fmt"

// Bank Card - Strategy

type Client struct {
	name         string
	age          int
	socialStatus Status
}

func (c *Client) pay() {
	c.socialStatus.discount()
}

func NewClient(name string, age int) *Client {
	if age > 50 {
		return &Client{name, age, PensionerStatus{}}
	}
	return &Client{name, age, DefaultStatus{}}
}

func (c *Client) setStatus(s Status) {
	c.socialStatus = s
}

type Status interface {
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
	newClient := NewClient("Olzhas", 20)
	newClient.pay()
	newClient.setStatus(StudentStatus{})
	newClient.pay()
}
