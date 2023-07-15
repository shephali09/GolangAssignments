package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter input string: ")
	fmt.Scanln(&input)
	containsA := strings.Contains(input, "a")
	containsE := strings.Contains(input, "e")
	containsP := strings.Contains(input, "p")

	if containsA && containsE && containsP {
		fmt.Println("All Present")
	} else if containsA || containsE || containsP {
		fmt.Println("One or more - Present")
	} else {
		fmt.Println("None Present")
	}
}
