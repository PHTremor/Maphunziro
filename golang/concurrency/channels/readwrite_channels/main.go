package main

import (
	"fmt"
	"time"
)

// receives the number of elements to iterate throgh the channel
func generator(ch chan int) {
	fmt.Println("generator fn awaits for n")
	// function will be blocked until the number is received
	// the more the sleep time in wait, the longer the blocking
	// initializing n with the value from the channel *Reading
	n := <-ch
	fmt.Println("n is", n)

	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + 1
	}

	// sending sum through the channel *Writing
	ch <- sum
}

func main() {
	// instatiating the channel *this is an unbuffered channel
	ch := make(chan int)

	// declaring a goroutine
	go generator(ch)

	fmt.Println("main awaits for result...")
	time.Sleep(time.Second)

	// sending value through the channel *Writing
	ch <- 5
	// *Reading -- getting value from the channel
	result := <-ch

	fmt.Println("result is: ", result)
}
