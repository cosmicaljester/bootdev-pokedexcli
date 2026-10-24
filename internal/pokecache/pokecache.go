package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mu       *sync.Mutex
}

func (c Cache) Add(key string, valye []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       valye,
	}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	elem, ok := c.entries[key]
	if ok {
		return elem.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		for key, elem := range c.entries {
			if time.Since(elem.createdAt) > c.interval {
				c.mu.Lock()
				delete(c.entries, key)
				c.mu.Unlock()
			}
		}
	}
}

func NewCache(inteval time.Duration) Cache {
	caccy := Cache{
		entries:  map[string]cacheEntry{},
		interval: inteval,
		mu:       &sync.Mutex{},
	}
	go caccy.reapLoop()
	return caccy
}
