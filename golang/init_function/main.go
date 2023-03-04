package main

import "fmt"

var x = xSetter()

func xSetter() int {
	fmt.Println("xSetter")
	return 42
}

// init functions are executed once per package by default when a program is run
// can be used fo (initial configurations) e.g initilize variables, detect OS, etc...
// ..without having the main package
// this is mostly used in libraries which normally do not have a main package
func init() {
	fmt.Println("init function")
}

func init() {
	fmt.Println("init-2 function")
}

func main() {
	fmt.Println("This is the main")
	fmt.Printf("x value is %d", x)
}

/* ## when a program is run and we import a package, the Go runtime executes in this order
1. initialize imported packages
2. Initialize and assign values to variables
3. run init functions


Run and check the order of the output
You can have many init functions
*/

/*
you can call init functions from another package without requiring to use that package inside your code
using the "import _ [package]" statement
	eg: import _"a" ==> in main
	package a has init function(s)
*/  