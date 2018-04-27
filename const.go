package goqrencode

/*
#include <qrencode.h>
#cgo LDFLAGS: -lqrencode
*/
import "C"

//QRencodeMode
const (
	QR_MODE_NUL        = C.QR_MODE_NUL        ///< Terminator (NUL character). Internal use only
	QR_MODE_NUM        = C.QR_MODE_NUM        ///< Numeric mode
	QR_MODE_AN         = C.QR_MODE_AN         ///< Alphabet-numeric mode
	QR_MODE_8          = C.QR_MODE_8          ///< 8-bit data mode
	QR_MODE_KANJI      = C.QR_MODE_KANJI      ///< Kanji (shift-jis) mode
	QR_MODE_STRUCTURE  = C.QR_MODE_STRUCTURE  ///< Internal use only
	QR_MODE_ECI        = C.QR_MODE_ECI        ///< ECI mode
	QR_MODE_FNC1FIRST  = C.QR_MODE_FNC1FIRST  ///< FNC1 first position
	QR_MODE_FNC1SECOND = C.QR_MODE_FNC1SECOND ///< FNC1 second position
)

//QRecLevel
const (
	QR_ECLEVEL_L = C.QR_ECLEVEL_L ///< lowest
	QR_ECLEVEL_M = C.QR_ECLEVEL_M
	QR_ECLEVEL_Q = C.QR_ECLEVEL_Q
	QR_ECLEVEL_H = C.QR_ECLEVEL_H ///< highest
)
