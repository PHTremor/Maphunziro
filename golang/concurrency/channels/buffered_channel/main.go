package main

import (
	"fmt"
	"time"
)

func Lolemba(ch chan string) {
	time.Sleep(time.Millisecond * 500)

	// Lolemba sending a message through the channel
	ch <- "This is Lolemba"
}

func Lachiwiri(ch chan string) {
	time.Sleep(time.Millisecond * 300)

	// Lachiwiri sending a message through the channel
	ch <- "This is Lachiwiri"
}

func main() {
	/*
		This is an unbuffered channel  where both sides have to ready to send and recieve data
		using this type of a channel will error bcoz the main func send a message through the channel
		when nobody is listening

		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
	*/
	// ch := make(chan string)

	/*
		This is a buffered channel where 1 is the buffer size (one-element buffer)
		Buffered channels can store values independently of the readiness of the
		sender and the receiver
		main func can send data through the channel when nobody is listeninng
	*/
	ch := make(chan string, 1)

	// Main sending a message through the channel
	// this is sent when nobody is listening
	ch <- "This is main"

	go Lolemba(ch)
	go Lachiwiri(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
