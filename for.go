package main

import "fmt"

func main() {
	
	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	slice := []string{"Desril", "Fatra", "Aditama", "Fahmi"}

	for i:=0; i<len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value:= range slice{
		fmt.Println("Index", i, "=" , value)
	}

	person := make(map[string]string)
	person["name"] = "Desril"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}