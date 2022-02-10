package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Smith"
	return
}

func main() {
	fistName, middleName, lastName := getFullName2()
	fmt.Println(fistName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}