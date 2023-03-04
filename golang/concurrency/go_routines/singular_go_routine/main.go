package main

import (
	"fmt"
	"time"
)

func ShowIt() {
	for {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Here it is!!")
	}
}

func main()  {
	// declaring a go routine
	// goroutines run concurrently with the main function and terminates when the main finishes
	go ShowIt()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Where is it?")
	}
}