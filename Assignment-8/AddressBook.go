package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Name            string
	MobileNumber    string
	AlternateMobile string
	Address         string
	City            string
}

func main() {

	addressBook := []Friend{
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
	file, err := os.Create("AddressBook.txt")
	if err != nil {
		fmt.Println("Error while creating the file:", err)
		return
	}
	defer file.Close()
	for _, friend := range addressBook {
		file.WriteString(fmt.Sprintf("Name: %s\n", friend.Name))
		file.WriteString(fmt.Sprintf("Mobile Number: %s\n", friend.MobileNumber))
		if friend.AlternateMobile != "" {
			file.WriteString(fmt.Sprintf("Alternate Mobile Number: %s\n", friend.AlternateMobile))
		}
		file.WriteString(fmt.Sprintf("City: %s\n", friend.City))
		file.WriteString("\n\n")
	}
	fmt.Println("data written to the file successfully!!")
}
