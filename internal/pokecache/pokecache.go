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
	lock  sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{cache: make(map[string]cacheEntry)}

	go newCache.reapLoop(interval)

	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		t := <-ticker.C
		c.lock.Lock()
		for key, val := range c.cache {
			if t.After(val.createdAt) {
				delete(c.cache, key)
			}
		}
		c.lock.Unlock()
	}
}
