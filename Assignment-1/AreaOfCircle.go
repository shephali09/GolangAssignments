package main

import "fmt"

func main() {
	var diameter float32 = 10
	var radius float32 = diameter / 2
	var areaOfCircle float32

	areaOfCircle = 3.14 * radius * radius
	fmt.Println("Area of Circle = ", areaOfCircle)

}
