package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("CricketPlayersStats.csv")
	if err != nil {
		fmt.Println("Error while opening the csv file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading csv file:", err)
		return
	}
	for _, record := range records {
		fmt.Println(record)
	}

	//Print all record except header row
	fmt.Printf("\nPrint all the records except header row:")
	for i, record := range records {
		if i > 0 {
			fmt.Println(record)
		}

	}
}
