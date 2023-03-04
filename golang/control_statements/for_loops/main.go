package main

import "fmt"

func main() {
	x := 5
	counter := x

	// First type of for loop
	for counter > 0 {
		fmt.Println("counter is: ", counter)

		counter--
	}

	// Second type of for loop
	for i := 0; i < x; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// Third type of for loop
	// Stops iterating if x is Even => break
	// jumps to the next iteration if x is odd => continue
	for {
		if x%2 != 0 {
			fmt.Printf("%d is an odd number\n", x)
			x++
			continue
		}
		break
	}

	// Forth type of for loop
	// an infinite loop
	// ...break => stops the loop after first iteration to avoid confusion
	for {
		fmt.Println("Never stops")
		break
	}

}
