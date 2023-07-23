package main

import "fmt"

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
	fmt.Println("Name and Mobile Numbers:")
	for _, friend := range addressBook {
		fmt.Printf("%s : %s\n", friend.Name, friend.MobileNumber)
	}

	fmt.Println("Names of Friends with Alternate Mobile Numbers:")
	for _, friend := range addressBook {
		if friend.AlternateMobile != "" {
			fmt.Println(friend.Name)
		}
	}

	cityFriendCount := make(map[string]int)
	for _, friend := range addressBook {
		cityFriendCount[friend.City]++
	}

	fmt.Println("Number of Friends in each city:")
	for city, count := range cityFriendCount {
		fmt.Printf("%s : %d\n", city, count)

	}

	fmt.Println("Names and Cities of College friends:")
	for _, friend := range addressBook {
		if friend.CollegeFriend {
			fmt.Printf("%s : %s\n", friend.Name, friend.City)
		}
	}
}
