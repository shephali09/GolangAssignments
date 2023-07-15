package main

import (
	"fmt"
	"time"
)

func main() {
	currentDateTime := time.Now()
	fmt.Println(currentDateTime)

	year := currentDateTime.Year()
	month := currentDateTime.Month()
	day := currentDateTime.Day()

	fmt.Println("Year:", year)
	fmt.Println("Month:", month)
	fmt.Println("Day:", day)

	fmt.Println(currentDateTime.Format("02-01-2006"))
}
