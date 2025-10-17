package main

import (
	"strings"
	"sync"
	"testing"
	"unsafe"
)

// go test -bench=. creation_test.go

func makeDirtySlice(size int) []byte {
	var sb strings.Builder
	sb.Grow(size)

	pointer := unsafe.StringData(sb.String())
	return unsafe.Slice(pointer, size)
}

var bytePool = sync.Pool{
	New: func() any {
		buf := make([]byte, 10<<20) // 10 MB заранее
		return &buf
	},
}

func makePooled(size int) []byte {
	buf := bytePool.Get().(*[]byte)
	// возвращаем слайс нужной длины (0, но с capacity)
	return (*buf)[:0:size]
}

func releasePooled(buf []byte) {
	bytePool.Put(&buf)
}

var Result []byte

func BenchmarkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = make([]byte, 0, 10<<20)
	}
}

func BenchmarkMakeDirty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = makeDirtySlice(10 << 20)
	}
}

func BenchmarkMakePooled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := makePooled(10 << 20)
		Result = buf
		releasePooled(buf)
	}
}
