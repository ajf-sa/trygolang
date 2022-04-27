package main

import "fmt"

func main() {
	fmt.Println(sliceSum([]int{1, 2, 3, 4, 5}))
}

func sliceSum(number_list []int) int {
	var sum int
	for _, number := range number_list {
		sum += number
	}
	return sum
}
