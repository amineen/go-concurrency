package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment(value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
	numGoroutines := 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter.Increment(2)
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.value)
}
