package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("LearnersData.txt")
	if err != nil {
		fmt.Println("Error while opening the txt file:", err)
		return
	}

	defer file.Close()

	fmt.Println("Read and Print Complete File content:")
	buffer := make([]byte, 128)
	var contain []byte
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Fatalf("Error while reading the txt file", err)
				return
			}
		}
		if n == 0 {
			break
		}
		contain = append(contain, buffer[:n]...)
	}
	fmt.Println(string(contain))

	//Read and print the data line by line
	fmt.Println("Reading one line at a time: ")
	file, err = os.Open("LearnersData.txt")
	if err != nil {
		fmt.Println("Error while opening the txt file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
