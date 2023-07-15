package main

import "fmt"

func MultiplesSum(number int) int {

	sum := 0
	for i := 1; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum

}

func main() {
	var number int
	fmt.Println("Enter number:")
	fmt.Scanln(&number)

	fmt.Println("Sum of multiples:", MultiplesSum(number))
}
