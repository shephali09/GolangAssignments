package main

import "fmt"

func average(numbers ...int) float64 {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}

	return float64(sum) / float64(len(numbers))
}
func main() {

	averageOfThree := average(50, 60, 30)
	fmt.Println("Average of 3 Integers:", averageOfThree)

	averageOfFive := average(10, 4, 5, 20, 12)
	fmt.Println("Average of 5 Integers:", averageOfFive)

}
