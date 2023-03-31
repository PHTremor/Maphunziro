package main

import (
	"fmt"
	"time"
)

/*
	ch :=make(chan int) //sender and receiver
	ch :=make(<- chan int) //receiver
	ch :=make(chan <- int) //sender
*/

// receiver can not send a message through the channel
func receiver(input <-chan int) {
	// loop throug the channel and print values until it closes
	for i := range input {
		fmt.Println(i)
	}
}

// sender can not receive a message on the channel
func sender(output chan<- int, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 500)
		// sending value through the channel
		output <- i * i
	}
	// closing the channel
	close(output)
}

func main() {
	// declaring the channel
	ch := make(chan int)

	// creating an instance of the   sender function -- can only write
	go sender(ch, 4)

	// creating an instance of the receiver function - can only read
	go receiver(ch)

	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
