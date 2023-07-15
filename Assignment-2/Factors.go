package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Println("Enter an integer:")
	fmt.Scanln(&number)

	if number%3 == 0 {
		fmt.Println("Pling")
	}

	if number%5 == 0 {
		fmt.Println("Plang")
	}

	if number%7 == 0 {
		fmt.Println("Plong")
	}

	if number%3 == 1 && number%5 == 1 && number%7 == 1 {
		fmt.Println(number)
	}

}
