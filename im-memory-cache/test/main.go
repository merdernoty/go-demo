package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache interface {
	Get(k string) (string bool)
	Set(k string, v string)
	cleanup(k string) bool
}

type CacheItem struct {
	value      string
	expiration int64
}

type InMemoryCache struct {
	data  map[string]CacheItem
	mutex sync.RWMutex
	ttl   time.Duration
}

func NewInMemoryCache(ttl time.Duration) *InMemoryCache {
	c := &InMemoryCache{
		data: make(map[string]CacheItem),
		ttl:  ttl,
	}
	go c.cleanup()
	return c
}

func (c *InMemoryCache) Get(k string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	data, ok := c.data[k]
	if !ok || time.Now().UnixNano() > data.expiration {
		return "", false
	}
	return data.value, ok
}

func (c *InMemoryCache) Set(k string, v string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[k] = CacheItem{
		value:      v,
		expiration: time.Now().Add(c.ttl).UnixNano(),
	}
}

func (c *InMemoryCache) cleanup() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now().UnixNano()
		c.mutex.Lock()
		for k, v := range c.data {
			if now > v.expiration {
				delete(c.data, k)
			}
		}
		c.mutex.Unlock()
	}
}

func main() {
	cache := NewInMemoryCache(5 * time.Second)

	cache.Set("foo", "bar")
	val, ok := cache.Get("foo")
	fmt.Println(val, ok) // bar true

	time.Sleep(6 * time.Second)

	val, ok = cache.Get("foo")
	fmt.Println(val, ok) // "" false
}
