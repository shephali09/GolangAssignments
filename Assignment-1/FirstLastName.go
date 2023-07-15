package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	fmt.Println("Enter First Name: ")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println(firstName,"", lastName)
}
