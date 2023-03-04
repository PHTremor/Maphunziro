package main

import (
	"fmt"
	"time"
)

/// #Channels provide a communication mechanism for concurrently running functions

// generator func runs in a goroutine that sums the first five integers and send it
// through the channel
func generator(ch chan int)  {
	sum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		sum = sum + 1
	}

	// sending through the channel
	ch <- sum
}

func main()  {
	// instantiating the channel
	ch := make(chan int)

	// declaring a goroutine
	go generator(ch)

	fmt.Println("main waits for results...")
	// wait until something is sent through the channel
	// and initialize result with the value
	result := <- ch

	fmt.Println(result)
}