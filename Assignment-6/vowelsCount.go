package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countVowels(input string) (int, map[rune]int) {
	count := 0
	vowels := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}

	input = strings.ToLower(input)

	for _, char := range input {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
			vowels[char]++
		}
	}
	return count, vowels

}

func main() {
	var input string
	fmt.Println("Enter String:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	count, vowelDistribution := countVowels(input)

	fmt.Println("Number of vowels:", count)
	fmt.Println("Vowel Distribution:")
	for vowel, count := range vowelDistribution {
		fmt.Printf("vowel['%c'] = %d\n", vowel, count)
	}
}
