package main

import (
	"bufio"
	"fmt"
	"os"
)

func encode(input string) []byte {
	encodedData := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		encodedData[i] = input[i] + 1
	}
	return encodedData
}

func decode(encoded []byte) string {
	decodedData := make([]byte, len(encoded))

	for i := 0; i < len(encoded); i++ {
		decodedData[i] = encoded[i] - 1
	}

	return string(decodedData)
}

func main() {
	var input string
	fmt.Println("Enter input string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	fmt.Println("Original string:", input)

	encoded := encode(input)
	fmt.Println("Encoded Bytes:", encoded)

	decoded := decode(encoded)
	fmt.Println("Decoded Data:", decoded)
}
