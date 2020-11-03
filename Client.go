package main

import "fmt"

type client struct {
	name         string
	age          int
	socialStatus status
}

func (c *client) setStatus(s status) {
	c.socialStatus = s
}

type Client struct {
	client
}

func (c *Client) NewClient(name string, age int) *Client {
	if age > 50 {
		return &Client{name}
	}
}

type status interface {
	pay()
}

type StudentStatus struct{}

func (s *StudentStatus) pay() {
	fmt.Println("Your discount is 20%")
}

type PensionerStatus struct{}

func (p *PensionerStatus) pay() {
	fmt.Println("Your discount is 30%")
}
