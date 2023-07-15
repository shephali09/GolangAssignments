package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var flip int
	flip = rand.Intn(2)
	if flip == 0 {
		fmt.Println("Head")
	} else {
		fmt.Println("Tail")
	}

}
