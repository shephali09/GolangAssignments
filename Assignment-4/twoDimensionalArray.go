package main

import "fmt"

func main() {
	inputArray := [5][2]int{
		{10, 0},
		{20, 0},
		{30, 0},
		{40, 0},
		{50, 0},
	}

	for i := 0; i < len(inputArray); i++ {
		inputArray[i][1] = inputArray[i][0] * 2
	}

	fmt.Println("Two dimensional array is:")
	for i := 0; i < len(inputArray); i++ {
		fmt.Println("{", inputArray[i][0], inputArray[i][1], "}")
	}
}
