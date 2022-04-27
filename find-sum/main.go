package main

import "fmt"

func main() {
	fmt.Println(findSum([]int{1, 2, 3, 4, 5}))
}

func findSum(number_list []int) int {
	var sum int
	for _, number := range number_list {
		sum += number
	}
	return sum
}
