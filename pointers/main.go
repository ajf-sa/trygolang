package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("Address of a is %x \n", &a)

	var ip *int = &a
	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	var ptr *int

	fmt.Printf("The value of ptr is : %x\n", ptr)
	fmt.Println(ptr == nil)
}
