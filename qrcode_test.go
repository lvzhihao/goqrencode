package goqrencode

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncodeString(t *testing.T) {
	qrcode, err := EncodeString("helowWorld", 0, QR_ECLEVEL_L, QR_MODE_8, 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", qrcode)
	t.Log(len(qrcode.data))
	offset := 0
	for y := 0; y < qrcode.width; y++ {
		for x := 0; x < qrcode.width; x++ {
			x := " "
			if int(qrcode.data[offset]&1) > 0 {
				x = "▉"
			}
			fmt.Print(" ", x)
			offset += 1
		}
		fmt.Println("")
	}
}

func TestEncodeInput(t *testing.T) {
	input, err := InputNew()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(input)
	err = input.Append(QR_MODE_8, bytes.NewReader([]byte("helowWorld")))
	if err != nil {
		t.Error(err)
		return
	}
	err = input.SetVersion(0)
	if err != nil {
		t.Error(err)
		return
	}
	qrcode, err := EncodeInput(input)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", qrcode)
	t.Log(len(qrcode.data))

}
