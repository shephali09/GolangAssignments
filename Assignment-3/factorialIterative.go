package main

import "fmt"

func factorialIterative(number int) int {
	factorial := 1
	for i := 2; i <= number; i++ {
		factorial *= i
	}
	return factorial

}
func main() {
	var number int
	fmt.Println("Enter number:")
	fmt.Scanln(&number)

	factorial := factorialIterative(number)
	fmt.Println("Factorial:", factorial)
}
