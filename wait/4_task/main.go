package main

import "sync"

type Data struct {
	mtx    sync.Mutex
	values []int
}

func (d *Data) Add(value int) {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	d.values = append(d.values, value)
}

func main() {
	data := Data{}
	data.Add(100)
}
