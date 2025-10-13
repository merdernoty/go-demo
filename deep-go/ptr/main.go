package main

import "fmt"

func main() {
	var value int32 = 100
	var ptr *int32 = &value

	fmt.Println("adress", ptr)
	fmt.Println("value", *ptr)

	*ptr = 500

	fmt.Println("adress", ptr)
	fmt.Println("value", *ptr)
}
