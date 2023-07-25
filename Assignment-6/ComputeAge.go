package main

import (
	"fmt"
	"time"
)

type Family struct {
	Name          string
	Relationsheip string
	DateOfBirth   string
}

func ComputeAge(DateOfBirth string) int {
	format := "2006-01-02"
	dob, err := time.Parse(format, DateOfBirth)
	if err != nil {
		fmt.Println("error while parsing", err)
		return -1
	}
	currentTime := time.Now()
	years := currentTime.Year() - dob.Year()

	if currentTime.Year() < dob.Year() {
		years--
	}
	return years
}

func main() {

	familyMembers := []Family{
		{Name: "Channaya", Relationsheip: "Father", DateOfBirth: "1958-06-01"},
		{Name: "Vidya", Relationsheip: "Mother", DateOfBirth: "1968-12-01"},
		{Name: "Harshad", Relationsheip: "Brother", DateOfBirth: "1992-08-15"},
	}

	for _, member := range familyMembers {
		age := ComputeAge(member.DateOfBirth)
		fmt.Printf("%s,%s Age:%d Years\n", member.Name, member.Relationsheip, age)
	}

}
