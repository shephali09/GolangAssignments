package main

import "fmt"

func FibonacciRecursion(number int) int {
	if number <= 1 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}
func main() {

	var number int
	fmt.Println("Enter the size of fibonacci series:")
	fmt.Scanln(&number)

	fmt.Println("Fibonacci Series is:")
	for i := 0; i < number; i++ {
		fmt.Print(FibonacciRecursion(i), " ")
	}
}
