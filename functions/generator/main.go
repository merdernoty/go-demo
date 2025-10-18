package main

import "fmt"

func Generator(number int) func() int {
	return func() int {
		r := number
		number++
		return r
	}
}

func main() {
	generator := Generator(1)
	for i := 0; i <= 10000000; i++ {
		fmt.Println(generator())
	}
}
