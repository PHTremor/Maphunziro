package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initializing the random generator
	rand.Seed(time.Now().UnixNano())
	// Generates a random number (0-1)
	x := rand.Float32()

	if x < 0.5 {
		fmt.Println("Head")
	} else {
		fmt.Println("Tail")
	}
}

// ## Emulates a coin toss...Head or Tail