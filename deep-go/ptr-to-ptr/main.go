package main

import "fmt"

func procces(temp **int32) {
	var value2 int32 = 200
	*temp = &value2
	fmt.Println(temp)
}

func main() {
	var value1 int32 = 100
	pointer := &value1

	fmt.Println("value", *pointer)
	fmt.Println("adress", pointer)
	procces(&pointer)

	fmt.Println("value", *pointer)
	fmt.Println("adress", pointer)
}

// 1var Copy-> prt1 Copy-> temp
