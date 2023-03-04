package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	A string
	B int32
	C string
}

func main() {
	t := T{"NiHao", 42, "Bye"}
	fmt.Println(t)

	// Elem returns the value that the pointer &t points to
	valueOfT := reflect.ValueOf(&t).Elem()

	// we'll loop through the struct T
	for i := 0; i < valueOfT.NumField(); i++ {
		f:= valueOfT.Field(i)

		// For each field of Type string, we'll set new value -- an uppercase of the value
		if f.Kind() == reflect.String {
			if f.CanSet() { //checking if the set operation is valid eg: for unexported fields
				current := f.String()
			// each Type has its own set method e.g: setInt32, setFloat32 etc
			f.SetString(strings.ToUpper(current))
			}
		}
	}

	// the modified values
	fmt.Println(t)
}