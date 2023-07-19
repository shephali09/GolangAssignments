package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	fmt.Println("Enter integers to add to the sorted slice. Enter 'X' to quit.")

	for {
		var input string
		fmt.Print("Enter an integer (or 'X' to quit): ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to quit.")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("Sorted slice:", slice)
	}
}
