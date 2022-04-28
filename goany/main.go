package main

import (
	"fmt"
	"reflect"
)

func main() {
	do(1)
	do("hello")
	do(true)
	do(struct{}{})
}

func do(v any) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println(reflect.TypeOf(v).Kind())
	}
}
