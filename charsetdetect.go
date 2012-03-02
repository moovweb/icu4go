package icu4go

// #cgo pkg-config: icu-i18n
//
// #include "helper.h"
import "C"
import "unsafe"
import "os"

type CharsetDetector struct {
	Ptr *C.UCharsetDetector
}

const EmptyEncoding = ""
const U_ZERO_ERROR = 0
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
	
	var status int
	var confidence int
	var guess string
	statusPtr := unsafe.Pointer(&status)
	confidencePtr := unsafe.Pointer(&confidence)
	detectorPtr := unsafe.Pointer(detector.Ptr)
	
	offset := 0
	maxConfidence := 0
	for ; offset < inputLen; offset += 8192 {
		inputPtr := unsafe.Pointer(&input[offset])
		guessPtr := C.detectCharset(detectorPtr, inputPtr, C.int(inputLen - offset), (*C.int)(confidencePtr), (*C.int)(statusPtr))
		if status == U_ZERO_ERROR {
			guess = C.GoString(guessPtr)
		}
		println(guess, confidence)
		if confidence > maxConfidence {
			bestGuess = guess
			maxConfidence = confidence
			if maxConfidence >= 100 {
				break
			}
		}
	}
	return
}


func (detector *CharsetDetector) Free() {
	if detector.Ptr != nil {
		C.ucsdet_close(detector.Ptr)
	}
}