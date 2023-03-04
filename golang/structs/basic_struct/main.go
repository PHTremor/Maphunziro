package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Rectangle struct {
	height int
	width  int
}

func main() {
	// instatiating the struc
	// h & w fields are initailized with 0
	a := Rectangle{}
	fmt.Println("a", a)

	// Or you can include the values as arguments
	b := Rectangle{4, 8}
	fmt.Println("b", b)

	c := Rectangle{width: 6, height: 2}
	fmt.Println("c", c)

	// height will be initialized with 0
	d := Rectangle{width: 4}
	fmt.Println("d", d)

	fmt.Println("====================================================\n checking for zero feilds")
	// There's are no rectangles with a height / width of zero
	// use a function to validate fields and create the struct like a constructor for the struct
	e, err := newRectangle(3, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("e", e)
	}

	fmt.Println("====================================================\n anonymous struct")
	// we can declare an anonymous struct like below
	// you can assign a known struct to an anonymous struct if they have the same fields
	// eg ac = a will not work because a has different fields from ac..(radius)
	ac := struct{x int; y int; radius int}{1,2,3}
	fmt.Printf("%+v\n", ac)
	fmt.Println("Type of a,b,c is", reflect.TypeOf(a))
	fmt.Println("Type of anonymous ac is ", reflect.TypeOf(ac))

}

// validates fields and creates the Rectangle struct
func newRectangle(height int, width int) (*Rectangle, error) {

	if height <= 0 || width <= 0 {
		return nil, errors.New("parameters must be greater than zero")
	}

	// returns a pointer to the struct
	return &Rectangle{height, width}, nil

}
