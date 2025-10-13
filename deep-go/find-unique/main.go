package main

import "fmt"

func main() {
	var arr = [...]int8{2, 3, 2, 3, 4, 4, 1}

	var number int8

	for _, element := range arr {
		fmt.Printf("%08b ^ %08b = &08b\n", number, element, number^element)
		number ^= element // XOR
	}
	fmt.Printf("\nResult: %08b\n", number)
}
