package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	data := [][]string{
		{"Name", "MobileNumber", "AlternateMobile", "Address", "City"},
		{"Shephali", "8329596787", "", "Sangli naka", "Kolhapur"},
		{"Aditi", "7798882067", "9420335206", "MG Road", "Kolhapur"},
		{"Neha", "8526541232", "", "Church Road", "Pune"},
		{"Aditya", "7412562356", "8598568574", "Main street", "Mumbai"},
		{"Harshad", "9595030156", "", "Narayan Peth", "Pune"},
		{"Onkar", "7894561236", "", "Sinhgad Road", "Pune"},
		{"Kaveri", "7412125262", "8548569871", "RK Nagar", "Solapur"},
		{"Parnika", "8459651235", "", "Borivali East", "Mumbai"},
		{"Aniket", "8529634563", "7459658523", "Sane Bridge", "Satara"},
		{"Akshata", "7894561230", "7050906084", "Wakad Bridge", "Pune"},
	}

	file, err := os.Create("AddressBook.csv")
	if err != nil {
		fmt.Println("Error while creating csv file:", err)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	for _, record := range data {
		err = writer.Write(record)
		if err != nil {
			fmt.Println("Error while writing data to csv file:", err)
			return
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			fmt.Println("Error while flushing the data:", err)
			return
		}

	}
	fmt.Println("File was written with all the data!")
}
