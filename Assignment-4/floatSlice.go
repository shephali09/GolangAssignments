/*Program takes floating point numbers from the user , till the user does not
enter q or Q. based on this input create a floating point slice.*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func GenerateFloatSlice(input string) []float64 {
	var numbers []float64
	for {
		fmt.Scanln(&input)

		if input == "q" || input == "Q" {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid floating-point number or 'q' to quit.")

		} else {
			numbers = append(numbers, num)
		}

	}
	return numbers
}

func average(numbers []float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum = sum + num
	}
	return sum / float64(len(numbers))

}

func standardDevitation(numbers []float64) float64 {
	avg := average(numbers)
	variance := 0.0
	for _, num := range numbers {
		variance += math.Pow(num-avg, 2)
	}
	variance /= float64(len(numbers))
	return math.Sqrt(variance)

}

func main() {
	var input string

	fmt.Println("Enter floating point number(Enter 'q' or 'Q' to quit):")
	numbers := GenerateFloatSlice(input)
	if len(numbers) == 0 {
		fmt.Println("No numbers entered")
		return
	}
	fmt.Println("Float slice:", numbers)

	average := average(numbers)
	fmt.Println("Average:", average)

	fmt.Println("Variance:", standardDevitation(numbers))
}
