package goqrencode

import (
	"bytes"
	"testing"
)

func TestVersionString(t *testing.T) {
	t.Log(ApiVersion())
	t.Log(ApiVersionString())
}

func TestInputCheck(t *testing.T) {
	err := InputCheck(QR_MODE_8, bytes.NewReader([]byte("helowWorld")))
	if err != nil {
		t.Error(err)
	}
}

func TestInput(t *testing.T) {
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
	t.Log(input.GetVersion())
}
