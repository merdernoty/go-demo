package main

import (
	"sync"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k string, v string)
}

type InMemoryCache struct {
	shards []Shard
}

func NewInMemoryCache(numShards int) *InMemoryCache {
	shards := make([]Shard, 0, numShards)
	for i := 0; i < numShards; i++ {
		shards = append(shards, Shard{data: make(map[string]string)})
	}
	return &InMemoryCache{
		shards: shards,
	}
}

func (s *InMemoryCache) Get(k string) (string, bool) {

}

func (s *InMemoryCache) Set(k string, v string) {

}

type Shard struct {
	data  map[string]string
	mutex sync.RWMutex
}

func (s *Shard) Get(k string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	data, ok := s.data[k]
	return data, ok
}

func (s *Shard) Set(k string, v string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[k] = v
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
