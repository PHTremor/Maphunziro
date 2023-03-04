package main

import "fmt"

type Coordinates struct {
	x int
	y int
}

type Circle struct {
	// Nested struct
	/* we also can have Coordinates as an embedded struct by not naming the field 'center' 
		ie...
		type Circle struct{
			Coordinates
			radius int
		}
	*/
	center Coordinates
	radius int
}

func main() {
	// Instatiation Circle requires that we instatiate  Coordinates
	c := Circle{Coordinates{1, 2}, 3}

	fmt.Printf("%+v\n", c)
}