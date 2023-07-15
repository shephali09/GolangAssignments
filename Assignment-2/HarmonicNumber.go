package main

import "fmt"

func main() {

	var number int
	fmt.Println("Enter a number:")
	fmt.Scanln(&number)
	var harmonicNumber float64

	for i := 1; i <= number; i++ {
		harmonicNumber += 1.0 / (float64)i
	}
	fmt.Println("The", number, "th Harmonic number is: ", harmonicNumber)
}
