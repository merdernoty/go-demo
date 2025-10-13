package main

import "fmt"

const (
	OpenModeIn     = 1 // 0000 0001
	OpenModeOut    = 2 // 0000 0010
	OpenModeAppend = 4 // 0000 0100
	OpenModeBinary = 8 // 0000 1000

	//Prepare sugar mask
	OpenModeIntOut = OpenModeIn | OpenModeOut // 0000 0001 + 0000 0010 = 0000 0011
)

func Open(filename string, mask int8) {
	if mask&OpenModeIn == 1 {
		fmt.Println("in mode")
	}
	if mask&OpenModeOut == 2 {
		fmt.Println("out mode")
	}
	if mask&OpenModeAppend == 4 {
		fmt.Println("append mode")
	}
	if mask&OpenModeBinary == 8 {
		fmt.Println("binary mode")
	}
}

func main() {
	Open("example.txt", OpenModeAppend|OpenModeBinary) // 0000 1100
	Open("example.txt", OpenModeOut|OpenModeIn)        // 0000 0011
	Open("example.txt", OpenModeAppend)                // 0000 0100
}
