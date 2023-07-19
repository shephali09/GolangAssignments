package main

import "fmt"

func isEven(num int) bool {
	return num%2 == 0
}

func isOdd(num int) bool {
	return num%2 == 1
}
func main() {
	originalArray := [20]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}
	startIndex := 4
	endIndex := startIndex + 10
	if endIndex > len(originalArray) {
		endIndex = len(originalArray)
	}

	slicedNumbers := originalArray[startIndex:endIndex]
	evenNumbers := make([]int, 0)
	oddNumbers := make([]int, 0)

	for _, num := range slicedNumbers {
		if isEven(num) {
			evenNumbers = append(evenNumbers, num)
		}

		if isOdd(num) {
			oddNumbers = append(oddNumbers, num)
		}
	}
	fmt.Println("Original Array:", originalArray)
	fmt.Println("Sliced Numbers:", slicedNumbers)
	fmt.Println("Even numbers:", evenNumbers)
	fmt.Println("Odd numbers:", oddNumbers)
}
