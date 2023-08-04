package main

import (
	"fmt"
	"strings"
)

func Map(inputSlice []string, condition func(string) string) []string {
	var result []string
	for _, word := range inputSlice {

		result = append(result, condition(word))

	}
	return result

}

func FirstCharacterCapital(word string) string {
	firstLetter := strings.ToUpper(string(word[0]))
	return firstLetter + word[1:]
}

func AddColonAtEnd(word string) string {
	return word + ":"
}

func main() {
	inputSlice := []string{"ball", "Apple", "cat", "butterfly", "kite", "Ant", "bark", "rat", "beat"}

	firstLetterCapital := Map(inputSlice, FirstCharacterCapital)
	fmt.Println("First letter of each word as Capital:", firstLetterCapital)

	colonAtEnd := Map(inputSlice, AddColonAtEnd)
	fmt.Println("Colon at the end of the word:", colonAtEnd)
}
