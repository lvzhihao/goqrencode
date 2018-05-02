package goqrencode

import (
	"image/color"
	"testing"
)

func TestImageDraw(t *testing.T) {
	qrcode, err := EncodeString("你好世界,helloWorld,こんにちは世界", 0, QR_ECLEVEL_L, QR_MODE_8, 1)
	if err != nil {
		t.Error(err)
		return
	}
	img := ImageNew()
	err = img.SetForegroundColor(color.RGBA{25, 111, 49, 1}).DrawQRcode(qrcode).WriteFile("/tmp/qrcode-111.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	err = img.SetForegroundColor(color.YCbCr{10, 10, 1}).DrawQRcode(qrcode).WriteFile("/tmp/qrcode-222.png")
	if err != nil {
		t.Error(err)
		return
	}
}
