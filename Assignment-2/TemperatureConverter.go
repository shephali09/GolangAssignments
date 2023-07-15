package main

import "fmt"

func main() {
	fmt.Println("Temperature Unit Converter")
	fmt.Println("0: Celsius to Fahrenheit")
	fmt.Println("1: Fahrenheit to Celsius")

	var choice int
	fmt.Println("Enter your choice o or 1:")
	fmt.Scanln(&choice)

	if choice == 0 {
		var celsius float64
		fmt.Println("Enter temperature in celsius: ")
		fmt.Scanln(&celsius)

		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Println("Temperature in Fahrenheit: ", fahrenheit)
	} else if choice == 1 {
		var fahrenheit float64
		fmt.Println("Enter temperature in fahrenheit: ")
		fmt.Scanln(&fahrenheit)

		celsius := (fahrenheit - 32) * 5 / 9
		fmt.Println("Temperature in Celsius:", celsius)

	} else {
		fmt.Println("Invalid inpur! Enter 0 or 1")
	}
}
