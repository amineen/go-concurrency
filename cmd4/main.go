package main

import (
	"fmt"
	"sync"
)

func main() {

	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing count:%d\n", count)
	}

	decrement := func(){
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing count:%d\n", count)
	}

	//increment count
	var arithmeticWg sync.WaitGroup
	for range 5{
		arithmeticWg.Add(1)
		go func(){
			defer arithmeticWg.Done()
			increment()
		}()
	}

	//decrement
	for range 5 {
		arithmeticWg.Add(1)
		go func(){
			defer arithmeticWg.Done()
			decrement()
		}()
	}
	arithmeticWg.Wait()
	fmt.Println("Arithmetic done")
}
