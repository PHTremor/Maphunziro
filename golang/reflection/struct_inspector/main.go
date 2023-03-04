package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Struct T
type T struct {
	B int
	C string
}

// Struct S
type S struct {
	C string
	D T
	E map[string]int
}

// function to print the reflected types
// receives an int(offset) && and a data Type
func printReflect(offset int, typeOfX reflect.Type) {
	// returns a new string consisting of count copies of the string [""]
	indent := strings.Repeat("", offset)
	fmt.Printf("%s %s: %s {\n", indent, typeOfX.Name(), typeOfX.Kind())

	// Checks if the data type is a Struct
	if typeOfX.Kind() == reflect.Struct {
		for i := 0; i < typeOfX.NumField(); i++ {
			innerIndent := strings.Repeat("", offset+4)
			f := typeOfX.Field(i)

			// checks the inner fields if they are of type struct...& prints with an indent if available
			if f.Type.Kind() == reflect.Struct {
				printReflect(offset+4, f.Type)
			} else {
				fmt.Printf("%s %s: %s\n", innerIndent, f.Name, f.Type)
			}
		}
	}

	fmt.Printf("%s }\n", indent)
}

func main() {
	// variable x of Type S
	x := S{
		"root",
		T{42, "forty Two"},
		make(map[string]int),
	}

	// function to inspect the struct
	printReflect(0, reflect.TypeOf(x))
}
