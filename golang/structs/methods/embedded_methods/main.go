package main

import (
	"fmt"
)

// #When you embedd a struct into another struct, its methods are also available
// #to the second struct... sort of like Inheritance

func main() {
	// instatiating Box...& Rectangle structs
	b := Box{Rectangle{3, 3}, 3}
	fmt.Printf("%+v\n", b)

	// calling Box's method
	fmt.Println("Volume", b.volume())
}

type Rectangle struct {
	height int
	width  int
}

// Method for Rectangle
func (r Rectangle) surface() int {
	return r.height * r.width
}

type Box struct {
	// embedded struct
	Rectangle
	depth int
}

// Method for Box
func (b Box) volume() int {
	// accesses Rectangle's method surface
	return b.surface() * b.depth
}
