package main

import "fmt"

func getHello(name string) {
	if name == ""{
		return "Hello Bro"
	}else{
		return "Hello " + name
	}
}

func main() {
	result := getHello("Desril")
	fmt.Println(result)

	fmt.Println(getHello(""))

	getHello("Budi")
}