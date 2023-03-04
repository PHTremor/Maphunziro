package main

import "fmt"

func main() {
	rect := Rectangle{2, 3}
	// any instance of Rectangle can invoke method 'surface'
	fmt.Printf("rectangle %v has surface %d", rect, rect.surface())

	fmt.Println("\n=============================== *Pointers")
	fmt.Println("rectangle", rect)

	rect.enlargeFalse(2)
	fmt.Println("rectangle withouPointer", rect)

	rect.enlargeTrue(2)
	fmt.Println("rectangle withPointer", rect)
}

type Rectangle struct {
	height int
	width  int
}

// this is a method: -- a special function with a receiver
// the receiver is (r Rectangle)
// The surface method calculates the surface area of a rectangle
func (r Rectangle) surface() int {
	return r.height * r.width
}

// You can also pass the arguments as pointers

// This method can't change the Rectangle values
// bcoz the receiver is a copy of Rectangle
func (r Rectangle) enlargeFalse(factor int) {
	r.height = r.height * factor
	r.width = r.width * factor
}

// This method is able to change Rectangle values
// bcoz the receiver is a pointer
func (r *Rectangle) enlargeTrue(factor int) {
	r.height = r.height * factor
	r.width = r.width * factor
}
