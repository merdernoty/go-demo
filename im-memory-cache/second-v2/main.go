package main

import (
	"sync"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k string, v string)
}

type InMemoryCache struct {
	data  map[string]string
	mutex sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]string),
	}
}

func (c *InMemoryCache) Get(k string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	data, ok := c.data[k]
	return data, ok
}

func (c *InMemoryCache) Set(k string, v string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[k] = v
}

func main() {
	cache := NewInMemoryCache()

	cache.Set("foo", "bae")
	cache.Set("baz", "qwx")
	cache.Set("gg", "sell")

	wg := sync.WaitGroup{}

	wg.Go(func() {
		cache.Set("foo", "updated_bar")
		println(1)
	})

	wg.Go(func() {
		cache.Set("baz", "updated_qwx")
		println(2)
	})

	wg.Go(func() {
		cache.Get("foo")
		println(3)
	})

	wg.Go(func() {
		cache.Get("baz")
		println(4)
	})

	wg.Wait()
}
