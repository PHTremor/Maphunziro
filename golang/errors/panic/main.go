package main

import "fmt"

func something() {
	defer fmt.Println("Closed Something")

	for i:=0;i<5;i++ {
		fmt.Println(i)

		if i > 2 {
			panic("Panic was called")
		}
	}
}

func main()  {
	defer fmt.Println("Closed main")
	something()
	fmt.Println("Something was finished")
}

// panic stops the execution flow, executes defered functions,
//  and returns control to main - the calling function
// The last print in main is never reached