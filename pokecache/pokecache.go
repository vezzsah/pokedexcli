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
	cachedData map[string]cacheEntry
	mu         *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	var cacheMap map[string]cacheEntry
	var mut sync.Mutex

	newCache := Cache{
		cachedData: cacheMap,
		mu:         &mut,
	}

	go newCache.reapLoop(interval)

	return &newCache
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var entry cacheEntry = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.cachedData[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cachedData[key]
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()

		for k, v := range c.cachedData {
			if time.Since(v.createdAt) >= interval {
				delete(c.cachedData, k)
			}
		}

		c.mu.Unlock()
	}

}
