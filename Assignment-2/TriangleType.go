package main

import "fmt"

func main() {
	var sideOne, sideTwo, sideThree float64
	fmt.Println("Enter the length of side 1:")
	fmt.Scanln(&sideOne)
	fmt.Println("Enter the length of side 2:")
	fmt.Scanln(&sideTwo)
	fmt.Println("Enter the length of side 3:")
	fmt.Scanln(&sideThree)

	if sideOne == sideTwo && sideTwo == sideThree && sideOne == sideThree {
		fmt.Println("Equilateral Triangle")
	} else if sideOne == sideTwo || sideTwo == sideThree || sideOne == sideThree {
		fmt.Println("Isosceles Triangle")
	} else {
		fmt.Println("Scalene Triangle")
	}
}
