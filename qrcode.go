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

/*
 * goqrencode QRcode
 */
type QRcode struct {
	code    *C.QRcode // libqrcode QRcode
	version int       // verion
	width   int       // width
	data    []byte    // origin
}

/*
 * new goqrencode qrcode
 */
func (c *QRcode) New(code *C.QRcode) *QRcode {
	c.code = code
	if c.code != nil {
		// fix goqrencode qrcode
		c.format()
		//set gc func
		runtime.SetFinalizer(c, (*QRcode).Destroy)
	}
	return c
}

/*
 * format params
 */
func (c *QRcode) format() *QRcode {
	if c.code != nil {
		c.version = int(c.code.version)
		c.width = int(c.code.width)
		c.data = C.GoBytes(unsafe.Pointer(c.code.data), C.int(c.width*c.width))
	}
	return c
}

/*
 * output matrix
 */
func (c *QRcode) Bitmap() (ret [][]bool) {
	ret = make([][]bool, c.width)
	offset := 0
	for row := 0; row < c.width; row++ {
		ret[row] = make([]bool, c.width)
		for i := 0; i < c.width; i++ {
			if int(c.data[offset]&1) > 0 {
				ret[row][i] = true
			} else {
				ret[row][i] = false
			}
			offset += 1
		}
	}
	return
}

/*
 * free libqrencode QRcode
 */
func (c *QRcode) Destroy() {
	if c.code != nil {
		C.QRcode_free(c.code)
	}
}

/*
 * Create a symbol from the string. The library automatically parses the input
 * string and encodes in a QR Code symbol.
 * see libqrencode QRcode_encodeString
 */
func EncodeString(str string, version int, level C.QRecLevel, mode C.QRencodeMode, casesensitive int) (*QRcode, error) {
	ret := C.QRcode_encodeString(C.CString(str), C.int(version), level, mode, C.int(casesensitive))
	if ret == nil {
		return nil, ThrowAndCatchException()
	} else {
		qrcode := new(QRcode)
		return qrcode.New(ret), nil
	}
}

/*
 * Create a symbol from the input data.
 * see libqrencode QRcode_encodeInput
 */
func EncodeInput(input *Input) (*QRcode, error) {
	ret := C.QRcode_encodeInput(input.input)
	if ret == nil {
		return nil, ThrowAndCatchException()
	} else {
		qrcode := new(QRcode)
		return qrcode.New(ret), nil
	}
}

//todo QRcodeList
