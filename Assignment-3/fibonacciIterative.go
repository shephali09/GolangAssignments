package main

import "fmt"

func FibonacciIterative(number int) []int {

	series := make([]int, number)
	series[0] = 0
	series[1] = 1

	for i := 2; i < number; i++ {
		series[i] = series[i-1] + series[i-2]
	}
	return series
}

func CountEven(series []int) int {
	count := 0
	for _, num := range series {
		if num%2 == 0 {
			count++
		}
	}
	return count
}
func main() {
	var number int
	fmt.Println("Enter the size of fibonacci series:")
	fmt.Scanln(&number)

	fibSeries := FibonacciIterative(number)
	evenCount := CountEven(fibSeries)
	fmt.Println("Fibonacci Series:", fibSeries)
	fmt.Println("Count of Even Numbers:", evenCount)
}
