package main

import "fmt"

type Mode struct {
	OpenModeIn     bool
	OpenModeOut    bool
	OpenModeAppend bool
	OpenModeBinary bool
}

var (
	OpenModeBinary = Mode{OpenModeIn: true, OpenModeBinary: true}
	OpenModeIntOut = Mode{OpenModeIn: true, OpenModeOut: true}
	OpenModeAppend = Mode{OpenModeAppend: true}
)

func Open(filename string, mode Mode) {
	if mode.OpenModeIn {
		fmt.Println("in mode")
	}
	if mode.OpenModeOut {
		fmt.Println("out mode")
	}
	if mode.OpenModeAppend {
		fmt.Println("append mode")
	}
	if mode.OpenModeBinary {
		fmt.Println("binary mode")
	}
}

func main() {
	Open("example.txt", OpenModeBinary)
	Open("example.txt", OpenModeIntOut)
	Open("example.txt", OpenModeAppend)
}
