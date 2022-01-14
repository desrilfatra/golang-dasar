package main

import "fmt"

func main() {
	const firstName string = "Desril"
	const lastName = "Fernandes"
	const value = 23

	fmt.Println("Nama Saya:", firstName, lastName)
	fmt.Println("Umur Saya:", value)

	const(
		firstName1 = "Rafif"
		lastName1 = "Jack"
		value1 = 25
	)
	fmt.Println("Nama Saya:", firstName1, lastName1)
	fmt.Println("Umur Saya:", value1)
	// error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"
}