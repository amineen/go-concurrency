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

	wg.Wait()
	fmt.Println("All goroutines complete.")

}
