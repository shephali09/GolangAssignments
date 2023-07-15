package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter number between 2 to 25: ")
	fmt.Scanln(&number)

	if number >= 2 && number <= 25 {
		for i := 1; i <= 10; i++ {
			result := number * i
			fmt.Println(result, "=", number, "*", i)
		}
	} else {
		fmt.Println("Invalid Input! Enter number between 2 to 25")
	}
}
