package main

import "fmt"

func countPlaceNames(placesArray []string) []int {
	characterCounts := make([]int, len(placesArray))

	for i, places := range placesArray {
		characterCounts[i] = len(places)
	}
	return characterCounts
}

func main() {
	numberOfPlaces := 7
	placesArray := make([]string, numberOfPlaces)

	fmt.Println("Enter the names of 7 places:")
	for i := 0; i < numberOfPlaces; i++ {
		fmt.Scan(&placesArray[i])
	}

	characterCountArray := countPlaceNames(placesArray)
	for i, place := range placesArray {
		fmt.Println(place, characterCountArray[i])
	}

}
