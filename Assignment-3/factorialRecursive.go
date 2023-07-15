package main

import "fmt"

func factorialRecusrsion(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorialRecusrsion(number-1)
}

func main() {
	var number int
	fmt.Println("Enter number:")
	fmt.Scanln(&number)

	factorial := factorialRecusrsion(number)
	fmt.Println("Factorial:", factorial)
}
