package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stateCapital := map[string]string{
		"Maharashtra":    "Mumbai",
		"Karanataka":     "Bengaluru",
		"Telengana":      "Hyderabad",
		"Rajasthan":      "Jaipur",
		"Goa":            "Panaji",
		"Tamil Nadu":     "Chennai",
		"Bihar":          "Patna",
		"Madhya Pradesh": "Bhopal",
		"West Bengal":    "Kolkata",
		"Uttar Pradesh":  "Lucknow",
	}
	fmt.Println(stateCapital)

	jsonEncoding, err := json.Marshal(stateCapital)
	if err != nil {
		fmt.Println("Error while converting map into JSON", err)
		return
	}

	fmt.Println("Conversion of stateCapital map into JSON data:", jsonEncoding)

}
