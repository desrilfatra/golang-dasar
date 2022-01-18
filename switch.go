package main

import "fmt"

func main() {

	name := "DesrilFatraAditamaRamadhan"

	switch name {
	case "Desril":
		fmt.Println("Hello Desril")
	case "Fatra":
		fmt.Println("Hello Fatra")
		fmt.Println("Hello Fatra")
	default:
		fmt.Println("Halo boleh Kenalan?")
	}

	// switch length := len(name); length >5 {
	// case true:
	// 	fmt.Println("Name is long")
	// case false:
	// 	fmt.Println("Name is short")
	// }

	length := len(name)
	switch {
		case length > 15:
			fmt.Println("Name is long")
		case length > 5:
			fmt.Println("Name is Cukup")
		default:
			fmt.Println("Name is short")
	}
}