package main

import (
	"fmt"
	"unsafe"
)

func accessToElements1() {
	data := [3]int{1, 2, 3}

	idx := 4

	fmt.Println(data[idx]) //panic
	fmt.Println(data[4])   //compilation error
}

func accessToElements2() {
	data := [3]int{1, 2, 3}

	idx := -1
	fmt.Println(data[idx]) //panic
	fmt.Println(data[-1])  // compilation error
}

func arrLen() {
	data := [10]int{}
	fmt.Println(len(data)) //10
}

func capLen() {
	data := [10]int{}
	fmt.Println(cap(data)) //10
}

func arrComparison() {
	first := [...]int{1, 2, 3}
	second := [...]int{1, 2, 3}

	fmt.Println(first == second)
	fmt.Println(first != second)
}

func emptyArr() {
	var data [10]byte

	fmt.Println(unsafe.Sizeof(data)) //10

	//data == nil  compilation error
}

func zeroArray() {
	var data [0]byte
	fmt.Println(unsafe.Sizeof(data)) //0
}

func zeroArrayStructs() {
	var data [10]struct{}
	fmt.Println(unsafe.Sizeof(data)) //0
}
