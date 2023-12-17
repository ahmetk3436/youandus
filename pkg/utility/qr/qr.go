package qr

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQrCode(url, fileName string) {
	err := qrcode.WriteFile(url, qrcode.Medium, 256, fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}
