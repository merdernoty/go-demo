package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Go(func() {
		fmt.Println("Привет из 1 горутины")
	})

	wg.Go(func() {
		fmt.Println("Привет их 2 горутины")
	})

	wg.Wait()
}
