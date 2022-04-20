package main

import (
	"image/jpeg"
	"os"

	"github.com/cristalhq/qrcode"
)

func main() {
	url := "hi"

	code, err := qrcode.Encode(url, qrcode.Q)
	checkErr(err)

	f, err := os.Create("qr.jpg")
	checkErr(err)
	defer f.Close()

	err = jpeg.Encode(f, code.Image(), nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
