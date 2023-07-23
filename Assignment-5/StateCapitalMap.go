package main

import (
	"fmt"
	"sort"
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

	var states []string
	for state := range stateCapital {
		states = append(states, state)
	}

	sort.Strings(states)

	fmt.Printf("Data Type of stateCapital: %T\n", stateCapital)
	fmt.Println("State and Capital Information with sorted states names:")
	for _, state := range states {
		fmt.Printf("%s → %s\n", state, stateCapital[state])
	}
}
