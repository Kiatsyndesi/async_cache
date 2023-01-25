package cache_v1

import "errors"

var ErrNotFound = errors.New("Value not found")

type CacheMethods interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Cache struct {
	storage map[string]string
}

func NewCache() CacheMethods {
	return &Cache{
		storage: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) error {
	c.storage[key] = value

	_, err := c.Get(key)
	if err != nil {
		return ErrNotFound
	}

	return nil
}

func (c *Cache) Get(key string) (string, error) {
	value, ok := c.storage[key]

	if !ok {
		return "", ErrNotFound
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
