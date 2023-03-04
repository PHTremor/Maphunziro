package main

import (
	"fmt"
	"reflect"
)

/*
	creating function at runtime
	Generics are currently not supported in Go, for instance reusing the same function for different
	data types
		eg: func Sum(a int, b int){...}
			func Sum(a float, b float){...}
	However using reflection --reflect.MakeFunc function-- we can build such a function that has
	different signatures
*/

func BuildAdder(i interface{})  {
	function := reflect.ValueOf(i).Elem()

	// MakeFunc receives a function Type and a function with a number of value Types inside an array
	// and returns an array of values
	newFunction := reflect.MakeFunc(function.Type(), func (in []reflect.Value)[]reflect.Value  {

		// we need not more than 2 arguments
		if len(in) > 2 {
			return []reflect.Value{}
		}

		a, b := in[0], in[1]

		// The arguments should be of the same type
		if a.Kind() != b.Kind() {
			return []reflect.Value{}
		}

		var result reflect.Value

		// The switch implements the addition according to the arguments Type
		switch a.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = reflect.ValueOf(a.Int() + b.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result = reflect.ValueOf(a.Uint() + b.Uint())	
		case  reflect.Float32, reflect.Float64:
			result = reflect.ValueOf(a.Float() + b.Float())
		case reflect.String:
			result = reflect.ValueOf(a.String() + b.String())
		default:
			result = reflect.ValueOf(interface{}(nil))
		}

		return []reflect.Value{result}
	})

	// setting the new function to the original function
	function.Set(newFunction)
}

func main()  {
	var intAdder func(int64, int64) int64
	var floatAdder func(float64, float64) float64
	var stringAdder func(string, string)string

	BuildAdder(&intAdder)
	BuildAdder(&floatAdder)
	BuildAdder(&stringAdder)

	fmt.Println(intAdder(1,2))
	fmt.Println(floatAdder(3.02, 2.433))
	fmt.Println(stringAdder("Nihao", "Go"))
}