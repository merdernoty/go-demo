package main

// go build -gcflags="-m" . | grep escape
func main() {
	var arrWithStack [10]int8
	_ = arrWithStack

	var arrWithHeap [10<<20 + 1]int8
	_ = arrWithHeap
}
