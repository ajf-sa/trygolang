package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()

}
