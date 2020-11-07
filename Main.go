// Group: SE1902
// Student 1: Ilyassov Olzhas
// Student 2: Sovetkazhiyev Alibek
// Student 3: Bakhytbekov Yersultan
// Task 1: Group Project
package main

import (
	"fmt"
	"os"
)

func main() {

	// Initialization of pizzas

	var name string
	var age int
	var isStudent bool
	fmt.Print("Enter Name: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Enter Age: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Print("Are you Student ?")
	fmt.Fscan(os.Stdin, &isStudent)

	fmt.Println(name, age, isStudent)

	var flag int

	for {
		fmt.Fscan(os.Stdin, &flag)
		if flag == 0 {
			break
		}
	}
}
