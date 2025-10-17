package main

import (
	"strings"
	"testing"
)

// go test -bench=. string-builder_test.go

func BenchmarkStringConcat(b *testing.B) {
	str := ""
	for i := 0; i < b.N; i++ {
		str += "test"
	}
	_ = str
}

func BenchmarkStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("test")
	}

	_ = builder.String()
}

func BenchmarkStringBuilderWithOptimization(b *testing.B) {
	builder := strings.Builder{}
	builder.Grow(4 + b.N*4)
	builder.WriteString("test")
	for i := 0; i < b.N; i++ {
		builder.WriteString("test")
	}

	_ = builder.String()
}
