package pokecache

import (
	"time"
	"mutex"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache map[string]cacheEntry {
	protect		mutex
	Cache		map[string]cacheEntry
}

func NewCache(interval time.Duration) {

}