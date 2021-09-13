package main

import "fmt"

func main() {
	// var balance [10]float32
	// var balance = [5]float32{1.2, 1.3, 2.5, 5.1, 6.2}
	var balance = []float32{1.2, 1.3, 2.5, 5.1, 6.2}
	balance[4] = 50.2
	var age float32 = balance[4]
	fmt.Println(balance)
	fmt.Println(age)

}
