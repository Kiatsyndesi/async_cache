package async

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiatsyndesi/async_cache/async/cache"
)

var ErrTimeout = errors.New("timeout")

type Cache struct {
	c *cache.Cache
}

func InitAsyncCache() *Cache {
	return &Cache{c: cache.InitCache()}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)

	go func() {
		defer close(ch)

		v, ok := c.c.Get(key)

		if ok {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return "", ErrTimeout
	case v, ok := <-ch:
		if ok {
			return v, nil
		}

		return "", errors.New(fmt.Sprintf("I can't find value for key %s", key))
	}
}

func (c *Cache) Add(ctx context.Context, key, value string) error {
	ch := make(chan any)

	go func() {
		defer close(ch)
		c.c.Add(key, value)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case <-ch:
		return nil
	}
}

func (c *Cache) Del(ctx context.Context, key string) error {
	ch := make(chan any)

	go func() {
		defer close(ch)
		c.c.Del(key)
	}()

	select {
	case <-ctx.Done():
		return ErrTimeout
	case <-ch:
		return nil
	}
}
