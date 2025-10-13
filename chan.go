package main

import (
	"fmt"
	"runtime"
)

// func writer() <-chan int {
// 	ch := make(chan int)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := range 5 {
// 			ch <- i + 1
// 		}

// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for i := range 5 {
// 			ch <- i + 11
// 		}

// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()
// 	return ch
// }

// func main() {
// 	ch := writer()

// 	for range 10 {
// 		v := <- ch
// 		fmt.Println("v =", v)
// 	}
// }

// func writer() <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for i := range 10 {
// 			ch <- i + 1
// 		}
// 	}()
// 	return ch
// }

// func doubler(input <-chan int) chan int {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for v := range input {
// 			time.Sleep(500 * time.Millisecond)
// 			ch <- v * 2
// 		}
// 	}()
// 	return ch
// }

// func reader(input <- chan int) {
// 	for v := range input{
// 		fmt.Println("value", v)
// 	}
// }

func main() {

	ch := make(chan int)

	go func() {
		for i := range 100000 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := range 100000 {
			ch <- i * 2
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("v =", v, "worker1")
			runtime.Gosched()
		}
	}()

	for v := range ch {
		fmt.Println("v =", v, "worker2")
	}
}
