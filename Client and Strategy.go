// Group: SE1902
// Student 1: Ilyassov Olzhas
// Student 2: Sovetkazhiyev Alibek
// Student 3: Bakhytbekov Yersultan
// Task 1: Group Project
package main

import "fmt"

type Client struct {
	fname        string
	name         string
	age          int
	socialStatus Status //Strategy
	//card Card //Bridge
}

//Factory
func NewClientFactory(socialStatus Status) func(name string, age int) *Client {
	return func(name string, age int) *Client {
		return &Client{name: name, age: age, socialStatus: socialStatus}
	}
}

func (c *Client) Pay() {
	c.socialStatus.Discount()
}

func (c *Client) SetStatus(s Status) {
	c.socialStatus = s
}

// Factory ?
func NewClient(name string, age int) *Client {
	if age > 50 {
		return &Client{name: name, age: age, socialStatus: PensionerStatus{}}
	}
	return &Client{name: name, age: age, socialStatus: DefaultStatus{}}
}

type Status interface {
	Discount()
}

type DefaultStatus struct{}

func (d DefaultStatus) Discount() {
	fmt.Println("Your discount is 0%")
}

type StudentStatus struct{}

func (s StudentStatus) Discount() {
	fmt.Println("Your discount is 10%")
}

type PensionerStatus struct{}

func (p PensionerStatus) Discount() {
	fmt.Println("Your discount is 20%")
}

func main() {

	newClient := NewClient("Olzhas", 20)
	newClient.Pay()
	newClient.SetStatus(StudentStatus{})
	newClient.Pay()

	/*

		oldClient := NewClientFactory(DefaultStatus{})
		//newClient := NewClientFactory(PensionerStatus{})

		card1 := &MasterCard{}
		//card2 := &Visa{}

		client1 := oldClient("Olzhas", card1,17)
		//client2 := newClient("Miras", card2,19)

		fmt.Println(client1)
		//fmt.Println(client2)

	*/
}
