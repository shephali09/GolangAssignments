package main

import (
	"fmt"
	"regexp"
)

func main() {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	var input string
	fmt.Println("Enter email-ID:")
	fmt.Scanln(&input)

	regexPatternCompiled, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error while compiling regex pattern", err)
		return
	}

	matches := regexPatternCompiled.MatchString(input)
	fmt.Println(matches)

}
