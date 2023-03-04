package main

import "fmt"

func something() {
	defer func() {
		r := recover()

		// resumes the execution
		if r != nil {
			fmt.Println("No need to panic if i=", r)
		}
	}()

	for i:=0; i< 5;i++ {
		fmt.Println(i)
		if i > 2 {
			// stops normal flow & runs defer functions
			panic(i)
		}
	}
	fmt.Println("Closed something normally")
}

func main()  {
	defer fmt.Println("Closed Main")

	something()
	fmt.Println("Main was finished")
}

// recover is used to resume normal execution after it has been stopped by
// panic