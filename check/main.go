package main

import "fmt"

func main() {
	a1 := make([]int, 0, 5)
	a1 = append(a1, []int{1, 2, 3, 4, 5,6,7,8,9,10,11}...)
	a1 = append(a1, 12,13)

	fmt.Println(a1, len(a1), cap(a1))
}
