package main

import "fmt"

func main() {
	var place1 string
	var place2 string
	var place3 string

	place1 = "Maldives"
	place2 = "Berlin"
	place3 = "Paris"

	//displays in same line
	fmt.Println(place1, ",", place2, ",", place3)

	//displays in new line
	fmt.Println(place1)
	fmt.Println(place2)
	fmt.Println(place3)
}
