package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int32
	B string
	c float32
}

func main() {
	t := T{42, "Forty Two", 8.72}

	valueOfT := reflect.ValueOf(t)
	// printing the structs' field values
	fmt.Println(valueOfT.Kind(), valueOfT)
	fmt.Println()

	// printing the value of each field in the struct
	for i := 0; i < valueOfT.NumField(); i++ {
		field := valueOfT.Field(i)

		fmt.Println(field.Kind(), field.String(), field)
		/*
			field.String returns a string like <int32 Value> for numeric values to inform you
			that there's an integer value in that fiels

			fmt.Println(field) prints the corresponding value as expected
		*/
		// fmt.Println(field)
	}
}
