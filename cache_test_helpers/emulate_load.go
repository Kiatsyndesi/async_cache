package cache_test_helpers

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var ErrNotFound = errors.New("Value not found")

type CacheMethods interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

func EmulateLoad(t *testing.T, c CacheMethods, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d - key", i)
		value := fmt.Sprintf("%d - value", i)

		wg.Add(1)

		//check writing to cache
		go func(key string) {
			err := c.Set(key, value)
			assert.NoError(t, err)

			wg.Done()
		}(key)

		//check reading from cache
		go func(key, value string) {
			storedValue, err := c.Get(key)

			if !errors.Is(err, ErrNotFound) {
				assert.EqualValues(t, value, storedValue)
			}

			wg.Done()
		}(key, value)

		//check for deleting
		go func(key string) {
			err := c.Delete(key)
			assert.NoError(t, err)

			_, err = c.Get(key)
			assert.ErrorIs(t, err, ErrNotFound)

			wg.Done()
		}(key)
	}

	wg.Wait()
}
