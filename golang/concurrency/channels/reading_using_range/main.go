package main

import "fmt"

func generator(ch chan int) {
	for i := 0; i < 5; i++ {
		// sending value in i through the channel
		ch <- i
	}
	// closing the channel
	close(ch)
}

func main() {
	// channel declaration
	ch := make(chan int)

	// instatiating the generator func
	go generator(ch)

	// consuming data using range
	// blocks execution and waits for values until the channel is closed
	for x := range ch {
		fmt.Println(x)
	}
	fmt.Println("Done")
}
