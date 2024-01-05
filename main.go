// greetings_with_age.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter your name: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Print("Enter your age: ")
	ageInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	age, err := strconv.Atoi(ageInput[:len(ageInput)-1])

	if err != nil {
		fmt.Println("Age reading error:", err)
		return
	}

	fmt.Printf("Hello %s! You are %d years old.", name[:len(name)-1], age)
}
