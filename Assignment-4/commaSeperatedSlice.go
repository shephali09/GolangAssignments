package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateSlice(input string) []int {
	var numSlice []int

	inputSlice := strings.Split(input, ",")
	numSlice = make([]int, len(inputSlice))
	for i, strNum := range inputSlice {
		numSlice[i], _ = strconv.Atoi(strNum)

	}
	return numSlice

}

func MaxumumAndNextHighest(numberSlice []int) (int, int) {

	max := numberSlice[0]
	nextHighest := numberSlice[0]

	for i := 1; i < len(numberSlice); i++ {
		if numberSlice[i] > max {
			nextHighest = max
			max = numberSlice[i]
		} else if numberSlice[i] > nextHighest && numberSlice[i] != max {
			nextHighest = numberSlice[i]
		}
	}
	return max, nextHighest
}

func MinimumAndNextLeast(numberSlice []int) (int, int) {
	min := numberSlice[0]
	nextLeast := numberSlice[0]

	for i := 1; i < len(numberSlice); i++ {
		if numberSlice[i] < min {
			nextLeast = min
			min = numberSlice[i]
		} else if numberSlice[i] < nextLeast && numberSlice[i] != min {
			nextLeast = numberSlice[i]
		}
	}

	return min, nextLeast
}

func main() {
	var input string
	fmt.Println("Enter comma-seperated numbers:")
	fmt.Scanln(&input)

	numberSlice := GenerateSlice(input)
	fmt.Println("Comma seperated input into slice:", numberSlice)

	maximumValue, nextHighestValue := MaxumumAndNextHighest(numberSlice)
	fmt.Println("Maximum:", maximumValue)
	fmt.Println("Next Highest Maximum:", nextHighestValue)

	minumumValue, nextLeastValue := MinimumAndNextLeast(numberSlice)
	fmt.Println("Minimum:", minumumValue)
	fmt.Println("Next Least Minimum:", nextLeastValue)
}
