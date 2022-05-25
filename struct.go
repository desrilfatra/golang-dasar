package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHai() {
	fmt.Println("Hello a", a.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Jakarta"
	eko.Age = 25

	eko.sayHi("Joko")
	eko.sayHai()

	// fmt.Println(eko)

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Bandung",
	// 	Age:     25,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Jakarta", 25}
	// fmt.Println(budi)
}
