package main

import (
	"fmt"
	"sync"
	"time"
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
	t0 := time.Now()
	var wg sync.WaitGroup
	counter := Counter{}
	numGoroutines := 10000
	wg.Add(numGoroutines)
	for range numGoroutines {
		go func() {
			defer wg.Done()
			counter.Increment(2)
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.value)

	// Define ANSI escape codes for colors
	green := "\033[32m"
	reset := "\033[0m"
	fmt.Printf("%scompleted in: %v%s\n", green, time.Since(t0), reset)

}
