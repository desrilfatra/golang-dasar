package main

import "fmt"

func main() {

	var ujian = 92
	var absensi = 74

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 75
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(ujian > 80 && absensi > 75)
}