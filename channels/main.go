package main

import (
	"fmt"
	"sync"
)

func main() {
	numberStream := make(chan int)

	go func() {
		defer close(numberStream)
		for i := range 2 {
			numberStream <- i
		}
	}()
	for x := range numberStream {
		fmt.Printf("Received: %d\n", x)
	}

	begin := make(chan any)
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", id)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	
	close(begin)
	wg.Wait()

}
