package main

import "fmt"

func main() {
	tos := []int{9, 87, 6, 5, 4, 3, 2, 1, 23}
	fmt.Println("count of slice:", sliceSum(tos))
	fmt.Println("index of:", binarySearch(tos, 23))

}
func sliceSum(number_list []int) int {
	var sum int
	for _, number := range number_list {
		sum += number
	}
	return sum
}

func binarySearch(slice []int, target int) int {
	start, end := 0, len(slice)-1
	for start <= end {
		mid := (start + end) / 2
		if target == slice[mid] {
			return mid
		} else if target < slice[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
