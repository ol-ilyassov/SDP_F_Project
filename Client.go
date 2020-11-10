// Group: SE1902
// Student 1: Ilyassov Olzhas
// Student 2: Sovetkazhiyev Alibek
// Student 3: Bakhytbekov Yersultan
// Task 1: Group Project
package main

// Visitor
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

type Client struct {
	name         string
	dateOfBirth  Date
	socialStatus Status //Strategy
}

func (c *Client) Pay() float32 {
	return c.socialStatus.Discount()
}

func (c *Client) SetStatus(s Status) {
	c.socialStatus = s
}

type Status interface {
	Discount() float32
}

type DefaultStatus struct{}

func (d DefaultStatus) Discount() float32 {
	fmt.Println("You have Your discount is 0%")
	return 0
}

type StudentStatus struct{}

func (s StudentStatus) Discount() float32 {
	fmt.Println("You have a Student Status.\n Consequently, your discount is 10%")
	return 10
}

type PensionerStatus struct{}

func (p PensionerStatus) Discount() float32 {
	fmt.Println("You have a Pensioner Status.\n Consequently, your discount is 20%")
	return 20
}

// Factory
func NewClientFactory(status Status) func(name string, dateOfBirth Date) *Client {
	return func(name string, dateOfBirth Date) *Client {
		return &Client{name, dateOfBirth, status}
	}
}

func NewClient(name string, date Date) *Client {
	StudentFactory := NewClientFactory(StudentStatus{})
	PensionerFactory := NewClientFactory(PensionerStatus{})
	DefaultFactory := NewClientFactory(DefaultStatus{})

	d1 := MakeDate(date.Year, date.Month, date.Day)
	d2 := time.Now().UTC()
	days := d2.Sub(d1).Hours() / 24
	age := math.Round(days / 365)
	fmt.Println(age)
	if age > 60 {
		return PensionerFactory(name, date)
	} else if (age >= 17) && (age <= 28) {
		fmt.Println("Are you Student? [Y/N]")
		var student string
		fmt.Fscan(os.Stdin, &student)
		if student == "Y" {
			return StudentFactory(name, date)
		}
	}
	return DefaultFactory(name, date)
}

func main() {
	date := Date{2001, 05, 30}
	client1 := NewClient("Olzhas", date)
	client1.Pay()

}
