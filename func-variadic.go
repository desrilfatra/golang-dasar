package main

import "fmt"

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Println(sum)

	slice := []int{1, 2, 3, 4, 5}
	sum = sumAll(slice...)
	fmt.Println(sum)
}