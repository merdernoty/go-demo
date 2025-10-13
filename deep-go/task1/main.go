package main

import (
	"fmt"
	"math"
)

func main() {
	var sined int8 = math.MaxInt8
	sined++

	var unsined uint8 = math.MaxUint8
	unsined++

	fmt.Println(sined)
	fmt.Println(unsined)

	var signed int8 = math.MaxInt8 + 1
	var unsigned uint8 = math.MaxUint8 + 1
}
