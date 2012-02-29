package icu4go

// #cgo pkg-config: icu-uc icu-i18n icu-io icu-le icu-lx
//
// #include "helper.h"
import "C"
import "unsafe"
import "os"

type CharsetDetector struct {
	Ptr *C.UCharsetDetector
}

const EmptyEncoding = ""
var U_ZERO_ERROR = int(C.U_ZERO_ERROR)
var ERR_CREATE_DETECTPR = os.NewError("cannot create charset detector")

func NewCharsetDetector() (detector *CharsetDetector, err os.Error) {
	detector = &CharsetDetector{}
	var status int
	statusPtr := unsafe.Pointer(&status)
	detector.Ptr = C.ucsdet_open((*C.UErrorCode)(statusPtr));
	if status != U_ZERO_ERROR {
		err = ERR_CREATE_DETECTPR
	}
	return
}

func (detector *CharsetDetector) GuessCharset(input []byte) (bestGuess string) {
	inputLen := len(input)
	if inputLen == 0 {
		return EmptyEncoding
	}
	inputPtr := unsafe.Pointer(&input[0])
	var status int
	statusPtr := unsafe.Pointer(&status)
	detectorPtr := unsafe.Pointer(detector.Ptr)
	
	bestGuessPtr := C.detectCharset(detectorPtr, inputPtr, C.int(inputLen), (*C.int)(statusPtr))
	if status == U_ZERO_ERROR {
		bestGuess = C.GoString(bestGuessPtr)
	}
	return
}


func (detector *CharsetDetector) Free() {
	if detector.Ptr != nil {
		C.ucsdet_close(detector.Ptr)
	}
}