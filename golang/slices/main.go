package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(a)
	fmt.Println(a[:])
	fmt.Println(a[0])
	fmt.Println(a[0],a[1],a[2])
	fmt.Println(a[0:2])
	fmt.Println(a[1:4])
	fmt.Println(a[:2])
	fmt.Println(a[2:])

	// to know the type of an object, use reflect
	fmt.Println(reflect.TypeOf(a))
	// a[0:3] is a slice hence it prints without its size
	fmt.Println(reflect.TypeOf(a[0:3]))
	fmt.Println(reflect.TypeOf(a[0]))

	// make and append functions
	// capacity is the difference btwn an array and a slice
	// cap is the momory reserved by the slice that hasnt been filled...the filled memory is the length
	fmt.Print("==================================================================== Make and Append")
	b := []int{0,1,2,3,4}
	fmt.Println(b, len(b), cap(b))

	// b is a slice for a
	c := append(b,5)
	fmt.Println(c, len(c), cap(c))

	c = append(c, 6)
	fmt.Println(c, len(c), cap(c))

	// a new slice from an existing one doesnt inherit the cap and len
	d := c[1:4]
	fmt.Println(d, len(d), cap(d))

	// make => creates a slice(Type.int,len.5,cap.10)
	e := make([]int,5,10)
	fmt.Println(e, len(e), cap(e))




}

// ## a slice is a reference to an array,
// ## it doesnt store data but provides a view of the data and provides access to the elements in the array

/*
ways to select a slice
	a[0] ==> Element @ position 0
	a[3:5] ==> from position 0 t0 5
	a[3:] ==> all elements from pos 3
	a[:3] ==> all elements from the start to pos 2
	a[:] ==> all elements
*/