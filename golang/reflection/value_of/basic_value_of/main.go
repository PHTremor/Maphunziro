package main

import (
	"fmt"
	"reflect"
)

/* reflect.Type can not access field values
reflect.Value does ...& and you can also access the Type too
*/

func main() {
	var a int32 = 42
	var b string = "Forty Two"

	valueOfA := reflect.ValueOf(a)
	fmt.Println(valueOfA.Interface())

	valueOfB := reflect.ValueOf(b)
	fmt.Println(valueOfB.Interface())

	// to know the type
	ValuePrint(a)
	ValuePrint(b)
}

// Checks the Type that's implementing the value
func ValuePrint(i interface{}) {
	v := reflect.ValueOf(i)

	// method Kind() returns a type -- which is a number representing the data Type
	// we use a switch to identify the type
	switch v.Kind() {
	case reflect.Int32:
		fmt.Println("Type: Int32, Value: ", v.Int())
	case reflect.String:
		fmt.Println("Type: String, Value: ", v.String())
	default:
		fmt.Println("Unknown Type")
	}
}
