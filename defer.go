package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Memulai Aplikasi")
	result := 10 / value
	fmt.Println("Hasilnya adalah", result)
}

func main() {
	runApplication(10)
}
