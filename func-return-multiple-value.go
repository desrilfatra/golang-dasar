package main

import "fmt"

func getFullName() (string,string,string){
	return "John","Doe","Smith"
}

func main(){
	firstname, _,lastname := getFullName()
	fmt.Println(firstname)
	// fmt.Println(middlename)
	fmt.Println(lastname)
}