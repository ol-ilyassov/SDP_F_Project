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
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Println(name, age)

	var flag int
	for {

		fmt.Fscan(os.Stdin, &flag)
		if flag == 0 {
			break
		}
	}
}
