package main

import (
	"fmt"
	"math/bits"
)

func SwapEndian[T ~uint16 | ~uint32 | ~uint64](value T) T {
	switch any(value).(type) {
	case uint16:
		return T(bits.ReverseBytes16(uint16(value)))
	case uint32:
		return T(bits.ReverseBytes32(uint32(value)))
	case uint64:
		return T(bits.ReverseBytes64(uint64(value)))
	default:
		panic("unsupported type")
	}
}

func swapEndianByte(value uint16) uint16 {
	return (value>>8)&0x00FF | (value<<8)&0xFF00
}

func main() {
	var (
		a uint16 = 0x1234
		b uint32 = 0x12319233
		c uint64 = 0x123456789ABCDEF0
	)
	fmt.Printf("16-bit: 0x%04X → 0x%04X\n", a, swapEndianByte(a))
	fmt.Printf("32-bit: 0x%08X → 0x%08X\n", b, SwapEndian(b))
	fmt.Printf("64-bit: 0x%016X → 0x%016X\n", c, SwapEndian(c))
}
