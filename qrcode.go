package goqrencode

/*
#include <qrencode.h>
#cgo LDFLAGS: -lqrencode
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type QRcode struct {
	code    *C.QRcode
	version int
	width   int
	data    []byte
}

func (c *QRcode) New() {
	runtime.SetFinalizer(c, (*QRcode).Destroy)
}

func (c *QRcode) Format() *QRcode {
	if c.code != nil {
		c.version = int(c.code.version)
		c.width = int(c.code.width)
		c.data = C.GoBytes(unsafe.Pointer(c.code.data), C.int(c.width*c.width))
	}
	return c
}

func (c *QRcode) Destroy() {
	if c.code != nil {
		C.QRcode_free(c.code)
	}
}

func EncodeString(str string, version int, level C.QRecLevel, mode C.QRencodeMode, casesensitive int) (*QRcode, error) {
	ret := C.QRcode_encodeString(C.CString(str), C.int(version), level, mode, C.int(casesensitive))
	if ret == nil {
		return nil, ThrowAndCatchException()
	} else {
		return &QRcode{
			code: ret,
		}, nil
	}
}
