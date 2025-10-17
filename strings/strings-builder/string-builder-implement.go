package main

import "fmt"

type Builder struct {
	buffer []byte
}

func NewBuilder() Builder {
	return Builder{}
}

func (b *Builder) Grow(cap int) {
	if cap < 0 {
		return
	}

	if cap < len(b.buffer) {
		b.buffer = b.buffer[:cap]
		return
	}

	buffer := make([]byte, len(b.buffer), cap)
	copy(buffer, b.buffer)
	b.buffer = buffer
}

func (b *Builder) Write(symbol byte) {
	b.buffer = append(b.buffer, symbol)
}

func (b *Builder) At(index int) *byte {
	if index < 0 || index >= len(b.buffer) {
		return nil
	}

	return &b.buffer[index]
}

func (b *Builder) String() string {
	return string(b.buffer)
}

func main() {
	builder := NewBuilder()

	builder.Grow(4)
	builder.Write('a')
	builder.Write('b')
	builder.Write('c')

	fmt.Println(builder.String())
}
