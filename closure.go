package main

import "fmt"

func main() {
	name := "desril"
	counter := 0

	increment := func() {
		name := "fatra"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
	
}