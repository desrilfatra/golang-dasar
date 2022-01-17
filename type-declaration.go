package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpDesril NoKtp = "123456789"
	var marriedStatus Married = true
	fmt.Println(noKtpDesril)
	fmt.Println(marriedStatus)
}