package main

import (
	"fmt"
	"math"
)

func Filter(numbers []float64, condition func(float64) bool) []float64 {
	var result []float64

	for _, num := range numbers {
		if condition(num) {
			result = append(result, num)
		}
	}
	return result
}

func NoDigitsAfterDecimalPoint(num float64) bool {
	return num == math.Trunc(num)

}

func DivisibleByThree(num float64) bool {
	return math.Mod(num, 3.0) == 0
}

func main() {
	numbers := []float64{5.12, 15.0, 3.3, 8.0, 9.0, 14.0, 6.63}

	noDigitsAfterDecimalPoint := Filter(numbers, NoDigitsAfterDecimalPoint)
	fmt.Println("Digits who is not having decimal part:", noDigitsAfterDecimalPoint)

	numbersDivisibleByThree := Filter(numbers, DivisibleByThree)
	fmt.Println("Numbers Divisible by Three:", numbersDivisibleByThree)

}
