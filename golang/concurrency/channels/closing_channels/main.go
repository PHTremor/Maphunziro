package main

import (
	"fmt"
	"time"
)

func sender(out chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		// writing|| sending i through the channel
		out <- i
	}
	// closing the channel
	close(out)
	fmt.Println("Sender Finished")
}

func main() {
	// creating an instance of the channel
	ch := make(chan int)

	// declaring a goroutine
	go sender(ch)

	// an endless loop
	for {
		// recieving values from the channel
		// num when the value is available; found is bool -- false when the channel is closed
		num, found := <-ch

		if found {
			fmt.Println(num)
		} else {
			fmt.Println("Finished: found is ", found)

			// break the endless loop
			break
		}
	}
}
