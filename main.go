package main

// import "fmt"

const (
	Read = 1 << iota
	Write
	Execute
	Delete
)

// func main() {
// 	fmt.Printf("Read=%d, Write=%d, Execute=%d Delete=%d\n", Read, Write, Execute, Delete)
// 	permissions := Read | Write | Execute
// 	if permissions&Read != 0 {
// 		fmt.Print("Есть права на чтение")
// 	}

// 	if permissions&Delete != 0 {
// 		fmt.Print("Есть права на удаление ")
// 	}
// }
