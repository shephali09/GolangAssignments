package main

import (
	"fmt"
	"strings"
)

func main() {
	var place string
	fmt.Println("Enter a place name where you woild like to visit most: ")
	fmt.Scanln(&place)

	upperCase := strings.ToUpper(place)
	fmt.Println(upperCase)

	lowerCase := strings.ToLower(place)
	fmt.Println(lowerCase)
}
