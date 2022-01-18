package main

import "fmt"
func main(){

	person := map[string]string{
		"name": "Desril",
		"address": "Jakarta",
		"age": "20",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"])


	var book map[string]string = make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "Desril"
	book["publisher"] = "Gramedia"
	book["year"] = "2020"
	book["ups"] = "wrong"
	
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}