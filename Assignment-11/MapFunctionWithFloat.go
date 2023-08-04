package main

import "fmt"

func Map(numbers []float64, condition func(float64) float64) []float64 {
	var result []float64
	for _, num := range numbers {
		result = append(result, condition(num))
	}
	return result
}

func ConvertPercentage(num float64) float64 {
	return num * 100
}

func HalfOfEachValue(num float64) float64 {
	return num / 2
}
func main() {
	numbers := []float64{0.6, 0.3, 0.88, 0.5, 0.65}

	convertIntoPercentage := Map(numbers, ConvertPercentage)
	fmt.Println("Convert each value as percentage:", convertIntoPercentage)

	halfOfTheValue := Map(numbers, HalfOfEachValue)
	fmt.Println("Half of each value:", halfOfTheValue)
}
