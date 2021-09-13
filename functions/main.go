package main

import "fmt"

func main() {
	cal := add(1, 2)
	fmt.Println(cal)
	addVoid(1, 2)

}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func addVoid(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}
