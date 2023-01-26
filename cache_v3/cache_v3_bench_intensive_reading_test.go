package cache_v3

import (
	"github.com/Kiatsyndesi/async_cache/cache_test_helpers"
	"testing"
)

const parallelFactor = 10_000

func Benchmark_Cache_With_Mutex (b *testing.B) {
	c := NewCache()

	for i := 0; i < b.N; i++ {
		cache_test_helpers.EmulateLoadForBenchWithIntensiveReading(c, parallelFactor)
	}
}
