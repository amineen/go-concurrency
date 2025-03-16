package main

import "fmt"

func main() {
	textStream := make(chan string)

	go func() {
		textStream <- "Hello World"
	}()

	fmt.Println(<-textStream)
}
