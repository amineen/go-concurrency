package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st Go routine sleeping...")
		time.Sleep(time.Nanosecond * 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd Go routine sleeping...")
		time.Sleep(time.Nanosecond * 2)
	}()

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	numGreeters := 8
	for id := range numGreeters {
		wg.Add(1)
		go hello(&wg, id)
	}

	wg.Wait()
	fmt.Println("All goroutines complete.")

}
