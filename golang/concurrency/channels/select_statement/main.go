package main

import (
	"fmt"
	"strconv"
	"time"
)

// sends numbers < 5 through the channel
func sendNumbers(out chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		out <- i
	}

	fmt.Println("no more numbers")
	// closing the channel
	close(out)
}

// sends messages through the channel
func sendMessages(out chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 300)

		out <- strconv.Itoa(i)
	}

	fmt.Println("no more messages")
	// closing the channel
	close(out)
}

func main() {
	// channel declaration
	numbers := make(chan int)
	messages := make(chan string)

	// creating instances of the functions
	go sendNumbers(numbers)
	go sendMessages(messages)

	// variables shall be used to know when a channel is closed
	closedNums, closedMsgs := false, false

	// The select stmt is executed only for a single opeartion
	// The For-loop is used to keep waiting for more messages
	for !closedNums || !closedMsgs {
		// select acts like switch stmt=> blocks execution until one operation is ready
		select {
		case num, ok := <-numbers: //reading from the numbers channel
			if ok {
				fmt.Printf("number %d\n", num)
			} else {
				// numbers channel is closed
				closedNums = true
			}

		case msg, ok := <-messages: //reading from the messages channel
			if ok {
				fmt.Printf("message %s\n", msg)
			} else {
				// messages channel is closed
				closedMsgs = true
			}

		}
	}

}
