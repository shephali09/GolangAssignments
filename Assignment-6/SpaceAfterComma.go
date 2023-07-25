package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func insertSpaceAfterComma(input string) string {
	parts := strings.Split(input, ",")

	correctedString := strings.Join(parts, ", ")
	return correctedString
}
func main() {
	var input string
	fmt.Println("Enter input string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	fmt.Println("Input:", input)

	output := insertSpaceAfterComma(input)
	fmt.Println("Output:", output)
}
