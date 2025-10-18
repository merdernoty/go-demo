package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Mul(x, y int) int {
	return x * y
}

func Calculate(x, y int, fn func(int, int) int) int {
	fmt.Printf("x=%d y=&d \n", x, y)
	return fn(x, y)
}

func CalcutateAdd(x, y int) int {
	fmt.Printf("x=%d y=&d\n", x, y)
	return Add(x, y)
}

func main() {
	Calculate(2, 3, Mul)
	CalcutateAdd(1, 2)
}
