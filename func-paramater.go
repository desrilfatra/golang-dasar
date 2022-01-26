package main

import "fmt"

func sayHellowTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main(){
	sayHellowTo("Desril", "Aditama")
}