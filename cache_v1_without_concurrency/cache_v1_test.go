package cache_v1_without_concurrency

import (
	"github.com/Kiatsyndesi/async_cache/cache_test_helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheV1(t *testing.T) {
	t.Parallel()

	testCache := NewCache()

	t.Run("Check stored value", func(t *testing.T) {
		t.Parallel()

		key := "Avenger"
		value := "Iron Man"

		err := testCache.Set(key, value)
		assert.NoError(t, err)

		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.EqualValues(t, value, storedValue)
	})

	t.Run("Check update of existing key", func(t *testing.T) {
		t.Parallel()

		key := "Avenger"
		value := "Iron Man"
		newValue := "Thor"

		err := testCache.Set(key, value)
		assert.NoError(t, err)

		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.EqualValues(t, value, storedValue)

		err = testCache.Set(key, newValue)
		assert.NoError(t, err)

		storedValue, err = testCache.Get(key)
		assert.NoError(t, err)

		assert.EqualValues(t, newValue, storedValue)
	})

	t.Run("Check for data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100000

		cache_test_helpers.EmulateLoad(t, testCache, parallelFactor)
	})
}
