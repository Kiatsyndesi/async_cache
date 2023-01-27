package async

import (
	"context"
	"testing"
	"time"
)

func Test_Async_Cache_Add(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond

	ctxBack := context.Background()
	ctxTimeout, _ := context.WithTimeout(ctxBack, to)

	err := ac.Add(ctxTimeout, "k", "v")
	if err != ErrTimeout {
		t.Error("expect for timeout")
	}

	to = time.Millisecond * 3
	ctxTimeout, _ = context.WithTimeout(ctxBack, to)

	err = ac.Add(ctxTimeout, "k", "v")
	if err != nil {
		t.Errorf("expect add %v", err)
	}
}

func Test_Async_Cache_Get(t *testing.T) {
	ac := InitAsyncCache()
	to := time.Millisecond

	key := "key"
	firstValue := "value"

	ctxBack := context.Background()
	ctxTimeout, _ := context.WithTimeout(ctxBack, to)

	_ = ac.Add(ctxBack, key, firstValue)
	v, err := ac.Get(ctxTimeout, key)

	if err != ErrTimeout {
		t.Error("expect for timeout")
	}

	ctxTimeout, _ = context.WithTimeout(ctxBack, to*5)
	v, err = ac.Get(ctxTimeout, key)
	if err != nil {
		t.Error("expect for get value")
	}

	if v != firstValue {
		t.Errorf("expect for %v, got %v", firstValue, v)
	}
}
