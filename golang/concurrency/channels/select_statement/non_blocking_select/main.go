package main

import "fmt"

// the select statement is blocking ie. the execution is blocked until one of the channels is ready
// the default case in select handles such situations
func main() {
	// this channel will never be ready
	ch := make(chan int)

	// this channel will be ready * uncomment and try
	// ch := make(chan int, 1)

	select {
	// reading from the channel
	case i := <-ch:
		fmt.Println("received", i)
	default:
		fmt.Println("Nothing received")
	}

	select {
	// sending through the channel
	case ch <- 42:
		fmt.Println("Sent 42")
	default:
		fmt.Println("Nothing sent")
	}

	select {
	// reading from the channel
	case i := <-ch:
		fmt.Println("received", i)
	default:
		fmt.Println("Nothing received")
	}
}
