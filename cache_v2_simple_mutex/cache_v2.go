package cache_v2_simple_mutex

import (
	"github.com/Kiatsyndesi/async_cache/cache_test_helpers"
	"sync"
)

type Cache struct {
	storage map[string]string
	mu      sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value

	/* doesn't work with mutex
	_, err := c.Get(key)
	if err != nil {
		return cache_test_helpers.ErrNotFound
	}
	*/
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.storage[key]

	if !ok {
		return "", cache_test_helpers.ErrNotFound
	}

	return value, nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)

	/* doesn't work with mutex
	_, err := c.Get(key)
	if err == nil {
		return errors.New("the key has not been removed\n")
	}
	*/
	return nil
}
