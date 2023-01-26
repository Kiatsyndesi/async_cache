package cache_test_helpers

import (
	"errors"
	"fmt"
	"sync"
)

func EmulateLoadForBenchWithIntensiveReading(c CacheMethods, parallelFactor int) {
	wg := &sync.WaitGroup{}

	for i := 0; i < parallelFactor / 10; i++ {
		key := fmt.Sprintf("%d - key", i)
		value := fmt.Sprintf("%d - value", i)

		wg.Add(1)
		//check writing to cache
		go func(k string) {
			err := c.Set(k, value)

			if err != nil {
				panic(err)
			}

			defer wg.Done()
		}(key)

		wg.Add(1)
		//check for deleting
		go func(k string) {
			err := c.Delete(k)

			if err != nil {
				panic(err)
			}

			defer wg.Done()
		}(key)
	}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d - key", i)
		value := fmt.Sprintf("%d - value", i)


		wg.Add(1)
		//check reading from cache
		go func(k, v string) {
			_, err := c.Get(k)

			if err != nil && !errors.Is(err, ErrNotFound) {
				panic(err)
			}

			defer wg.Done()
		}(key, value)
	}

	defer wg.Wait()
}
