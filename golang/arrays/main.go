package main

import "fmt"

func main() {
	// declaring an empty array of size 5
	// empty arrays are filled with zeros
	var a [5]int
	fmt.Println(a)

	// another way to declare an array
	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println(b)

	// empty indexes are filled with zeros
	c := [5]int{0, 1, 2}
	fmt.Println(c)

	// array length
	fmt.Println(len(a))

	// iterating an array using for
	fmt.Print("=========================== Loops")
	names := []string{
		"Frank", "Tremor", "Katswiri",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	// using range
	// range returns the index and the corresponding item
	for position, name := range names {
		fmt.Println(position, name)

		// the returned item from range is a copy of the item
		/*
			this will not work
				name = name + "_changed"
		*/
		// to modify the actual item acess the original variable using the item's index
		names[position] = name + "_changed"

	}
	fmt.Println(names)
}

// ## arrays have a fixed size
