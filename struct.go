package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var Desril Customer
	Desril.Name = "Desril"
	Desril.Address = "Tangsel"
	Desril.Age = 23

	fmt.Println(Desril.Name)
	fmt.Println(Desril.Address)
	fmt.Println(Desril.Age)
}
