package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	//Returns a function literal (clousure)
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	addPng := func(name string) string { return name + ".png" }
	addJpg := MakeAddSuffix(".jpg")
	fmt.Println(addPng("filename"), addJpg("filename"))
	//output filename.png filename.jpg
}
