package main

import "fmt"

func main(){
	var temp_in_degree float64
	var temp_in_fahrenheit float64

	fmt.Println("Enter Temperature in degree c: ")
	fmt.Scanln(&temp_in_degree)

	temp_in_fahrenheit = (temp_in_degree * 9/5) +32
	
	fmt.Println("Temperature in Fahrenheit : ",temp_in_fahrenheit)

}