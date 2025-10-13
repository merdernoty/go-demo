package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(HasDublicat(arr))
}

func HasDublicat(data []int) bool {
	var lookup int8
	for _, num := range data {
		if lookup&(1<<num) != 0 {
			return true
		}

		lookup |= 1 << num
	}
	return false

}
