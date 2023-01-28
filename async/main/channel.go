package main

import (
	"github.com/Kiatsyndesi/async_cache/async/cache"
	"github.com/Kiatsyndesi/async_cache/cache_test_helpers"
)

type Data struct {
	Key   string
	Value string
}

type ChanCache struct {
	c      *cache.Cache
	InChan chan Data
}

func InitChanCache() *ChanCache {
	return &ChanCache{
		c:      cache.InitCache(),
		InChan: make(chan Data),
	}
}

func (c *ChanCache) GetChannel() chan<- Data {
	return c.InChan
}

func (c *ChanCache) Run() {
	go func() {
		for {
			v, ok := <-c.InChan

			if !ok {
				return
			}

			c.c.Add(v.Key, v.Value)
		}
	}()
}

func (c *ChanCache) Get(key string) (string, error) {
	v, ok := c.c.Get(key)

	if !ok {
		return "", cache_test_helpers.ErrNotFound
	}

	return v, nil
}

func (c *ChanCache) ChannelGet(key string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		v, ok := c.c.Get(key)

		if ok {
			ch <- v
		}
	}()

	return ch
}
