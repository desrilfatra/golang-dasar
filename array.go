package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Desril"
	names[1] = "Fatra"
	names[2] = "Aditama"
	names[3] = "Fahmi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [4]int{
		1,
		100,
		200,
		300,
		
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])

	fmt.Println(len(names))
	fmt.Println(len(values))
	
}