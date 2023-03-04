package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// init the random generator
	rand.Seed(time.Now().UnixNano())
	id := rand.Int() %6

	artist, err := getMusician(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[%d] %s", id, artist)
	}
}

// an Array of type string
var Musicians = []string{
	"Athume", "Mandela", "Tremor", "Hayze",
}

// retuns a musician or an error
func getMusician(id int) (string, error) {
	if id < 0 || id >= len(Musicians) {
		return "", errors.New(
			fmt.Sprintf("invalid id [%d]", id))
	}

	return Musicians[id], nil
}
