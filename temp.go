package main

/*
// Factory
func NewClient(name string, date Date) *Client {
	d1 := MakeDate(date.Year, date.Month, date.Day)
	d2 := time.Now().UTC()
	days := d2.Sub(d1).Hours() / 24
	age := math.Round(days / 365)
	fmt.Println(age)
	if age > 60 {
		return &Client{name, date, PensionerStatus{}}
	} else if (age >= 17) && (age <= 28) {
		fmt.Println("Are you Student? [Y/N]")
		var student string
		fmt.Fscan(os.Stdin, &student)
		if student == "Y" {
			return &Client{name, date, StudentStatus{}}
		}
	}
	return &Client{name, date, DefaultStatus{}}
}
*/

/*
func main() {
	visa1 := &Visa{}
	mastercard1 := &MasterCard{}

	order1 := &OrdinaryOrder{}

	order1.SetCard(visa1)
	order1.Pay()
	fmt.Println()

	order1.SetCard(mastercard1)
	order1.Pay()
	fmt.Println()
}

*/
