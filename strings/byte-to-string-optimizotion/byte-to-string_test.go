package main

import (
	"testing"
	"unsafe"
)

//go test -bench=. -benchmem byte-to-string_test.go

func Convert(data []byte) string {
	if ledn(data) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(data), len(data))
}

var Result string

func BenchmarkConvert(b *testing.B) {
	slice := []byte("Hello World")
	for i := 0; i < b.N; i++ {
		Result = string(slice)
	}
}

func BenchmarkConvertOptimization(b *testing.B) {
	slice := []byte("Hello World")
	for i := 0; i < b.N; i++ {
		Result = Convert(slice)
	}
}
