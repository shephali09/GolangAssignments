package main

import "fmt"

func Sum(a float64, b float64) float64 {
	return a + b
}

func Difference(a float64, b float64) float64 {
	return a - b
}

func getSumAndFifference(a float64, b float64) (float64, float64) {
	return Sum(a, b), Difference(a, b)
}
func main() {
	var a, b float64
	fmt.Println("Enter a:")
	fmt.Scanln(&a)

	fmt.Println("Enter b:")
	fmt.Scanln(&b)

	sumResult, differenceResult := getSumAndFifference(a, b)

	fmt.Println("Sum and Difference : ", sumResult)
	fmt.Println("Sum and Difference : ", differenceResult)

}
