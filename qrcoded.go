package main

import (
	"fmt"
	"io/ioutil"
)

func GenerateQRCode(code string) []byte {
	return []byte{0xFF}
}

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}
