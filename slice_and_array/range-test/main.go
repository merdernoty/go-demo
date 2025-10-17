package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3}
	// for i, v := range arr {
	// 	arr = append(arr, 2)
	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)
	// }

	for i := 0; i < len(arr); i++ {
		arr = append(arr, 2)
		fmt.Printf("Index: %d, Value: %d\n", i, arr[i])
	}
}
