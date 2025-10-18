package main

func readNilMap() {
	var data map[int]int
	_ = data[100]
}

func deleteNilMap() {
	var data map[int]int
	delete(data, 100)
}

func writeNilMap() {
	var data map[int]int
	data[100] = 1
}

func rangeNilMap() {
	var data map[int]int
	for range data {
	}
}

func rewriteExistKey() {
	data := make(map[int]int)
	data[100] = 100
	data[100] = 1000
}

func deleteNotExistKey() {
	data := make(map[int]int)
	delete(data, 100)
}

func main() {
	deleteNotExistKey()
}
