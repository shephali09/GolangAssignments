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

	//convert json data into map
	err := json.Unmarshal([]byte(jsonData), &recipeMap)
	if err != nil {
		fmt.Println("Error while Unmarshaling data:", err)
		return
	}

	for key, value := range recipeMap {
		fmt.Printf("%s : %s\n", key, value)
	}

	//converting data to json and writing to the json file
	file, err := os.Create("ApplePieRecipe.json")
	if err != nil {
		fmt.Println("Error while creating json file:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(recipeMap)
	if err != nil {
		fmt.Println("Error while writing the data:", err)
		return
	}
	fmt.Println("JSON file written successfully")

}
