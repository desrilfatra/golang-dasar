package main

import "fmt"

func main() {

	var name1 = "Desril"
	var name2 = "Desrilf"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 102
	var value2 = 200

	fmt.Println(value1>value2)
	fmt.Println(value1<value2)
	fmt.Println(value1==value2)
	fmt.Println(value1!=value2)
}