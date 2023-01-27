package cache_v1_without_concurrency

import (
	"errors"
	"github.com/Kiatsyndesi/async_cache/cache_test_helpers"
)

type Cache struct {
	storage map[string]string
}

func NewCache() cache_test_helpers.CacheMethods {
	return &Cache{
		storage: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) error {
	c.storage[key] = value

	_, err := c.Get(key)
	if err != nil {
		return cache_test_helpers.ErrNotFound
	}

	return nil
}

func (c *Cache) Get(key string) (string, error) {
	value, ok := c.storage[key]

	if !ok {
		return "", cache_test_helpers.ErrNotFound
	}

	return value, nil
}

func (c *Cache) Delete(key string) error {
	delete(c.storage, key)

	_, err := c.Get(key)
	if err == nil {
		return errors.New("the key has not been removed\n")
	}

	return nil
}
