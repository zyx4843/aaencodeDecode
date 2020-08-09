package main

import (
	"io/ioutil"
	"os"

	"github.com/zyx4843/aaencodeDecode"
)

func main() {
	data, err := ioutil.ReadFile("index.js")
	if err == nil {
		encStr := aaencodeDecode.Aaencode(string(data))
		ioutil.WriteFile("encode.js", []byte(encStr), os.ModePerm)
		decStr := aaencodeDecode.Aadecode(encStr)
		ioutil.WriteFile("decode.js", []byte(decStr), os.ModePerm)
	}
}
