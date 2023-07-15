package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountVowelConsonant(inputString string) (int, int) {

	vowelCount := 0
	consonantCount := 0

	for _, char := range inputString {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			vowelCount++
		} else {
			consonantCount++

		}
	}
	return vowelCount, consonantCount
}

func main() {
	var inputString string
	fmt.Println("Enter string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString = scanner.Text()

	}

	vowelCount, consonantCount := CountVowelConsonant(inputString)
	fmt.Println("Vowel Count:", vowelCount)
	fmt.Println("Consonant Count:", consonantCount)
}
