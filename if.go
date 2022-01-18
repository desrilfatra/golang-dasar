package main

import "fmt"

func main() {
	name := "Aditama"

	if name == "Desril" {
		fmt.Println("Hello Desril")
	}else if name == "Fatra" {
		fmt.Println("Hello Fatra")
	}else {
		fmt.Println("Halo boleh Kenalan?")
	}

	
	if length := len(name); length > 6 {
		fmt.Println("Name is long")
	}else {
		fmt.Println("Name is short")
	}
		
	
}

