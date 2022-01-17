package main

import "fmt"

func main() {
	var result = 10 + 5
	fmt.Println(result)


	var a = 10
	var b = 5
	var c = a % b
	fmt.Println(c)

	// assignment
	var d = 10
	d += 5
	fmt.Println(d)

	// unary
	var e = 10
	e++
	fmt.Println(e)
}