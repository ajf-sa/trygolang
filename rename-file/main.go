package main

import (
	"log"
	"os"
)

func main() {
	orginalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(orginalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
