package main

import "fmt"

func main() {
	numberStream := make(chan int)

	go func() {
		defer close(numberStream)
		for i := range 5 {
			numberStream <- i
		}
	}()
	for x := range numberStream {
		fmt.Printf("Received: %d\n", x)
	}
}
