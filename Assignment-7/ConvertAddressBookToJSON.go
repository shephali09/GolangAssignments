package main

import (
	"encoding/json"
	"fmt"
)

type Friend struct {
	Name            string
	MobileNumber    string
	AlternateMobile string
	Address         string
	City            string
	CollegeFriend   bool
}

func main() {

	addressBook := []Friend{
		{"Shephali", "8329596787", "", "Sangli naka", "Kolhapur", false},
		{"Aditi", "7798882067", "9420335206", "MG Road", "Kolhapur", true},
		{"Neha", "8526541232", "", "Church Road", "Pune", true},
		{"Aditya", "7412562356", "8598568574", "Main street", "Mumbai", false},
		{"Harshad", "9595030156", "", "Narayan Peth", "Pune", true},
		{"Onkar", "7894561236", "", "Sinhgad Road", "Pune", true},
		{"Kaveri", "7412125262", "8548569871", "RK Nagar", "Solapur", true},
		{"Parnika", "8459651235", "", "Borivali East", "Mumbai", true},
		{"Aniket", "8529634563", "7459658523", "Sane Bridge", "Satara", false},
		{"Akshata", "7894561230", "7050906084", "Wakad Bridge", "Pune", false},
	}
	fmt.Println(addressBook)

	fmt.Printf("\nConversion of slice of struct into json:")
	jsonEncoding, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Println("Error while converting into JSON", err)
		return
	}

	fmt.Println(jsonEncoding)

	fmt.Printf("\nConversion of JSON data to original slice of struct:")
	var ConvertedAddressBookOne []Friend
	err = json.Unmarshal(jsonEncoding, &ConvertedAddressBookOne)
	if err != nil {
		fmt.Println("Error while converting JSON data to slice of struct", err)
		return
	}
	fmt.Println(ConvertedAddressBookOne)

}
