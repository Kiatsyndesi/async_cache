package cache_test_helpers

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func EmulateLoad(t *testing.T, c CacheMethods, parallelFactor int) {
	wg := &sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d - key", i)
		value := fmt.Sprintf("%d - value", i)

		wg.Add(1)
		//check writing to cache
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)

			defer wg.Done()
		}(key)

		wg.Add(1)
		//check reading from cache
		go func(k, v string) {
			storedValue, err := c.Get(k)

			if !errors.Is(err, ErrNotFound) {
				assert.Equal(t, v, storedValue)
			}

			defer wg.Done()
		}(key, value)

		wg.Add(1)
		//check for deleting
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			/*
				_, err = c.Get(key)
				assert.ErrorIs(t, err, ErrNotFound)
			*/
			defer wg.Done()
		}(key)
	}

	defer wg.Wait()
}
