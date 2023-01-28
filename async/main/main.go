package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	cc := InitChanCache()
	wg.Add(1)

	go func() {
		defer wg.Done()
		ch := cc.GetChannel()

		for i := 0; i < 100; i++ {
			ch <- Data{Key: fmt.Sprintf("%d - key", i), Value: fmt.Sprintf("%d - value", i)}
			time.Sleep(time.Millisecond)
		}
	}()
}
