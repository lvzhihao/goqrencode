package goqrencode

import (
	"fmt"
	"testing"

	"github.com/lvzhihao/goutils"
)

func TestEncodeString(t *testing.T) {
	qrcode, err := EncodeString("https://github.com/lvzhihao/gozbar/blob/24c7533dd6f3f366a71d95905110dbdcc6b5ec38/symbol.go", 5, QR_ECLEVEL_L, QR_MODE_8, 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", qrcode.Format())
	t.Log(len(qrcode.data))
	offset := 0
	for y := 0; y < qrcode.width; y++ {
		for x := 0; x < qrcode.width; x++ {
			x := " "
			if goutils.ToString(qrcode.data[offset]&1) == "1" {
				x = "x"
			}
			fmt.Print(" ", x)
			offset += 1
		}
		fmt.Println("")
	}
}
