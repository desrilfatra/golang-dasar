package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHellow(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (p Person) getName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) getName() string {
	return a.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	sayHellow(eko)

	cat := Animal{
		Name: "Kucing",
	}
	sayHellow(cat)
}
