package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	fmt.Println("Application Barcode")

	// Create Barcode
	qrCode, _ := qr.Encode("www.goole.com", qr.M, qr.Auto)

	// Scale the barcode here
	qrCode, _ = barcode.Scale(qrCode, 300, 300)

	// Create output file
	file, _ := os.Create("file/qrcode.png")

	// Encode the barcode to png
	png.Encode(file, qrCode)
}
