package goqrencode

/*
#include <qrencode.h>
#cgo LDFLAGS: -lqrencode
*/
import "C"
import (
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"unsafe"
)

func ThrowAndCatchException() error {
	// todo
	// patch libqrencode
	// extern "C"
	// int throw_and_catch_exception {
	//     try {
	//        ....
	//     } catch (...) {
	//        return ...
	//     }
	// }
	return fmt.Errorf("libqrencode error")
}

/*
 * Validate the input data.
 * See qrencode.h QRinput_check
 */
func InputCheck(mode C.QRencodeMode, r interface{}) error {
	b, err := ioutil.ReadAll(r.(io.Reader))
	if err != nil {
		return err
	}
	ret := C.QRinput_check(mode, C.int(len(b)), (*C.uchar)(unsafe.Pointer(&b)))
	if int(ret) == 0 {
		return nil
	} else {
		return ThrowAndCatchException()
	}
}

/*
 * Return a string that identifies the library version.
 */
func ApiVersion() (major, minor, micro int) {
	C.QRcode_APIVersion((*C.int)(unsafe.Pointer(&major)), (*C.int)(unsafe.Pointer(&minor)), (*C.int)(unsafe.Pointer(&micro)))
	return
}

/*
 * Return a string that identifies the library version.
 */
func ApiVersionString() string {
	return C.GoString(C.QRcode_APIVersionString())
}

/*
 * goqrencode Input
 */
type Input struct {
	input *C.QRinput
}

/*
 * create goqrencode Input
 */
func InputNew() (ret *Input, err error) {
	ret = new(Input)
	err = ret.New()
	return
}

/*
 * create Input
 */
func (c *Input) New() error {
	c.input = C.QRinput_new()
	if c.input == nil {
		return ThrowAndCatchException()
	} else {
		runtime.SetFinalizer(c, (*Input).Destroy)
		return nil
	}
}

/*
 * Append data to an input object
 * The data is copied and appended to the input object.
 * See qrencode.h QRinput_append
 *
 */
func (c *Input) Append(mode C.QRencodeMode, r interface{}) error {
	b, err := ioutil.ReadAll(r.(io.Reader))
	if err != nil {
		return err
	}
	ret := C.QRinput_append(c.input, mode, C.int(len(b)), (*C.uchar)(unsafe.Pointer(&b)))
	if int(ret) == 0 {
		return nil
	} else {
		return ThrowAndCatchException()
	}
}

/*
 * Append ECI header.
 * See qrcnecode.h QRinput_appendECIheader
 */
func (c *Input) AppendECIHeader(ecinum uint) error {
	ret := C.QRinput_appendECIheader(c.input, C.uint(ecinum))
	if int(ret) == 0 {
		return nil
	} else {
		return ThrowAndCatchException()
	}
}

/*
 * Get current version.
 * See qrcnecode.h QRinput_getVersion
 */
func (c *Input) GetVersion() int {
	return int(C.QRinput_getVersion(c.input))
}

/*
 * Set version of the QR code that is to be encoded.
 * See qrcnecode.h QRinput_setVersion
 */
func (c *Input) SetVersion(version int) error {
	ret := C.QRinput_setVersion(c.input, C.int(version))
	if int(ret) == 0 {
		return nil
	} else {
		return ThrowAndCatchException()
	}
}

/*
 * Get current error correction level.
 * See qrcnecode.h QRinput_getErrorCorrectionLevel
 */
func (c *Input) GetErrorCorrectionLevel() C.QRecLevel {
	return C.QRinput_getErrorCorrectionLevel(c.input)
}

/*
 * Set error correction level of the QR code that is to be encoded.
 * See qrcnecode.h QRinput_setErrorCorrectionLevel
 */
func (c *Input) SetErrorCorrectionLevel(level C.QRecLevel) error {
	ret := C.QRinput_setErrorCorrectionLevel(c.input, level)
	if int(ret) == 0 {
		return nil
	} else {
		return ThrowAndCatchException()
	}
}

/*
 * Set version and error correction level of the QR code at once.
 * See qrcnecode.h QRinput_setVersionAndErrorCorrectionLevel
 */
func (c *Input) SetVersionAndErrorCorrectionLevel(version int, level C.QRecLevel) error {
	ret := C.QRinput_setVersionAndErrorCorrectionLevel(c.input, C.int(version), level)
	if int(ret) == 0 {
		return nil
	} else {
		return ThrowAndCatchException()
	}
}

/*
 * free input object
 * See qrcnecode.h QRinput_free
 */
func (c *Input) Destroy() {
	C.QRinput_free(c.input)
}
