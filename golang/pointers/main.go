package main

import "fmt"

func main() {
	// The Value will not be changed, bcoz we did'nt pass the pointer to the variable x
	x := 100
	a(x)
	fmt.Println(x)

	// changes value bcoz of the pointer passed:
	// func b receives the pointer to the variable passed
	// '&' returns the the pointer to the variable
	b(&x)
	fmt.Println(x)
	// prints the memory location
	fmt.Println(&x)

}

// has no pointer
func a(i int) {
	i = 0
}

// For Type T, *T indicates a pointer to a value of that type...T
// *int is a pointer to the value in i *int
func b(i *int) {
	*i = 0
}

// ## Pass Pointers and not just the value if the value has to be modified
// ## Pass just the value if it does'nt have to be changed
