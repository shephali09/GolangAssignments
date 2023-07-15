package main

import "fmt"

func main() {
	var number int
	var TenTimesNumber int
	fmt.Println("Enter number: ")
	fmt.Scanln(&number)
	TenTimesNumber = (number * 10)
	fmt.Println("Ten times of", number, "is:", TenTimesNumber)
}
