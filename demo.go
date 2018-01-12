package main

import (
	"fmt"

	"github.com/satbirdd/cgo_xunfei_example/iat"
)

func main() {
	fileName := "/home/liulei/works/demo/bin/wav/iflytek02.wav"

	output := iat.Iat(fileName)

	fmt.Printf("XUNFEI output: %v", output)
}
