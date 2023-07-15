package main

import (
	"fmt"
	"math"
)

func main() {
	var principalAmount, year, rateOfInterest float64
	fmt.Println("Enter the Principal amount:")
	fmt.Scanln(&principalAmount)
	fmt.Println("Enter the number of years:")
	fmt.Scanln(&year)
	fmt.Println("Enter the rate of interest:")
	fmt.Scanln(&rateOfInterest)

	r := rateOfInterest / (12 * 100)
	n := year * 12

	payment := principalAmount * r / (1 - math.Pow(1+r, n) - 1)

	fmt.Println("The EMI for a loan:", payment)

}
