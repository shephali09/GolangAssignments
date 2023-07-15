package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var RandomNumberOne, RandomNumberTwo, RandomNumberThree, RandomNumberFour, RandomNumberFive int
	var average float64

	RandomNumberOne = rand.Intn(40) + 10
	RandomNumberTwo = rand.Intn(40) + 10
	RandomNumberThree = rand.Intn(40) + 10
	RandomNumberFour = rand.Intn(40) + 10
	RandomNumberFive = rand.Intn(40) + 10

	fmt.Println(RandomNumberOne, RandomNumberTwo, RandomNumberThree, RandomNumberFour, RandomNumberFive)
	average = float64(RandomNumberOne+RandomNumberTwo+RandomNumberThree+RandomNumberFour+RandomNumberFive) / 5
	fmt.Println(average)
}
