package main

import (
	"fmt"
	"math/rand"
)

func numbersInRange(numbers []int, min, max int) int {
	count := 0
	for _, num := range numbers {
		if num > min && num <= max {
			count++
		}
	}
	return count
}

func main() {

	min := 0
	max := 25
	var numbers []int
	count := 25
	for i := 0; i < count; i++ {
		randomNumber := rand.Intn(max-min) + min
		numbers = append(numbers, randomNumber)
	}

	buckets := [5][2]int{
		{0, 5},
		{6, 10},
		{11, 15},
		{16, 20},
		{21, 25},
	}

	numberDistribution := make(map[string]int)
	for _, bucket := range buckets {
		key := fmt.Sprintf("> %d && <= %d", bucket[0], bucket[1])
		numberDistribution[key] = numbersInRange(numbers, bucket[0], bucket[1])
	}

	fmt.Println("Generated numbers:", numbers)
	fmt.Println("Number distribution:")
	for key, count := range numberDistribution {
		fmt.Printf("%s: %d\n", key, count)
	}

}
