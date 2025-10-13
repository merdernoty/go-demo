package main

type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

type InMemoryCache struct {
	data map[string]string
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]string),
	}
}

func (c *InMemoryCache) Set(k string, v string) {
	c.data[k] = v
}

func (c *InMemoryCache) Get(k string) (string, bool) {
	data, ok := c.data[k]
	return data, ok
}

func main() {
	cache := NewInMemoryCache()
	cache.Set("foo", "bar")
	data, ok := cache.Get("foo")
	if ok {
		println(data)
		return
	}
	println("Key not found")
}
