package main

import (
	"fmt"
	"strings"
)

func Filter(inputSlice []string, condition func(string) bool) []string {
	var result []string
	for _, word := range inputSlice {
		if condition(word) {
			result = append(result, word)
		}
	}
	return result
}

func WordsStartWithB(word string) bool {
	return strings.HasPrefix(strings.ToLower(word), "b")
}

func HasThreeLetters(word string) bool {
	return len(word) == 3
}

func main() {
	inputSlice := []string{"Ball", "Apple", "Cat", "Butterfly", "Kite", "Ant", "Bark", "Rat", "Beat"}

	wordStartWithB := Filter(inputSlice, WordsStartWithB)
	fmt.Println("Words Starts With B:", wordStartWithB)

	threeLetters := Filter(inputSlice, HasThreeLetters)
	fmt.Println("Words which have 3 letters:", threeLetters)
}
