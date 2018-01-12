package iat_sample

/*
#cgo CFLAGS: -I/home/liulei/works/demo/include
#cgo LDFLAGS: -L/home/liulei/works/demo/libs/x64 -lmsc -lrt -ldl -lpthread
#include "./iat_sample_2.c"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Iat(fileName string) string {
	// fileName := "/home/liulei/works/demo/bin/wav/iflytek02.wav"
	cName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cName))

	output := C.Entry(cName)

	return C.GoString(output)
}
