package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := `{
		"Thinly Sliced Peeled Apples" : "6 Cups",
        "sugar" : "3/4 cup",
        "flour" : "2 tablespoons",
        "cinamon" : "1/4 teaspoon",
        "nutmeg" : "1/8 tablespoon",
        "lemon juice": "1 tablespoon"
	}`

	var recipeMap map[string]string

	err := json.Unmarshal([]byte(jsonData), &recipeMap)
	if err != nil {
		fmt.Println("Error while JSON data:", err)
		return
	}

	for key, value := range recipeMap {
		fmt.Printf("%s : %s\n", key, value)
	}

	file, err := os.Create("ApplePie.txt")
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	for key, value := range recipeMap {
		file.WriteString(fmt.Sprintf("%s : %s\n", key, value))
	}
	fmt.Println("Dara written to the successfully!!")
}
