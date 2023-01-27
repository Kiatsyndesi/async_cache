package cache_test_helpers

import "errors"

var ErrNotFound = errors.New("value not found")

type CacheMethods interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}
