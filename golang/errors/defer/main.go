package main

import "fmt"

func closedMsg() {
	fmt.Println("Closed!!!")
}

func main()  {
	defer closedMsg()

	fmt.Println("Doing Something...")
	defer fmt.Println("Certainly Closed!!!")
	fmt.Println("Doing Something else...")
}

// ## defer executes after the function end
// #* defer function are executed in inverse order
// ## useful to clean up resources allocated to a function