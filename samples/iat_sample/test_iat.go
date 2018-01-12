package main

import (
	"github.com/satbirdd/cgo_xunfei_example/samples/iat_sample"
)

func main() {
	fileName := "/home/liulei/works/demo/bin/wav/iflytek02.wav"

	output := iat_sample.Iat(fileName)

	fmt.Printf("-----------> oupt: %v", C.output)
}
