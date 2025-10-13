package main

import "fmt"

func main() {
	fmt.Println(PowerOfTwo(128))
}

func PowerOfTwo(value int) bool {
	return value > 0 && value&(value-1) == 0
}
