package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func MakeDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

type baseClient interface {
	GetName() string
	GetStatus() string
	MakeDiscount() float32
}

type Client struct {
	name         string
	dateOfBirth  Date
	socialStatus Status //Strategy DP
}

func (c *Client) GetName() string {
	return c.name
}
func (c *Client) GetStatus() string {
	return c.socialStatus.String()
}
func (c *Client) MakeDiscount() float32 {
	return c.socialStatus.Discount()
}

// Strategy DP
type Status interface {
	Discount() float32
	String() string
}
type DefaultStatus struct{}

func (d DefaultStatus) Discount() float32 {
	fmt.Println(" - Your discount is 0% - ")
	return 0
}
func (d DefaultStatus) String() string {
	return " - You have an Account with Default status - "
}

type StudentStatus struct{}

func (s StudentStatus) Discount() float32 {
	fmt.Println(" - Your discount is 10% - ")
	return 10
}
func (s StudentStatus) String() string {
	return " - You have an Account with Student status - "
}

type PensionerStatus struct{}

func (p PensionerStatus) Discount() float32 {
	fmt.Println(" - Your discount is 20% - ")
	return 20
}
func (p PensionerStatus) String() string {
	return " - You have an Account with Pensioner status - "
}

// Factory DP
func NewClientFactory(name string, date Date) *Client {
	d1 := MakeDate(date.Year, date.Month, date.Day)
	d2 := time.Now().UTC()
	days := d2.Sub(d1).Hours() / 24
	age := math.Round(days / 365)
	if age > 60 {
		return &Client{name, date, PensionerStatus{}}
	} else if (age >= 17) && (age <= 28) {
		fmt.Print("> Are you Student? [Y/N]: ")
		var student string
		fmt.Fscan(os.Stdin, &student)
		if student == "Y" {
			return &Client{name, date, StudentStatus{}}
		}
	}
	return &Client{name, date, DefaultStatus{}}
}
