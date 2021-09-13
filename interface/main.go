package main

import "fmt"

type Result interface {
	Count() int
}

type Add struct {
	num1, num2 int
}

func (a Add) Count() int {
	return a.num1 + a.num2
}

type Mult struct {
	num1, num2 int
}

func (a Mult) Count() int {
	return a.num1 * a.num2
}
func (a Mult) ToString() int {
	return a.num1 * a.num2
}

func getCount(result Result) int {
	return result.Count()
}

func main() {
	add := Add{num1: 4, num2: 5}
	mul := Mult{num1: 4, num2: 6}
	fmt.Println(getCount(add))
	fmt.Println(getCount(mul))
}
