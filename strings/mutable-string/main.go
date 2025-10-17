package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := []byte("Hello World")
	strData := unsafe.String(unsafe.SliceData(str), len(str))

	fmt.Println(strData)
	str[0] = 'w'
	fmt.Println(strData)
}
