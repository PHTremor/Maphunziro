package main

// # Reflection == expore, introspect, and modify code during run-time
import (
	"fmt"
	"reflect"
)

func main() {
	var a int32 = 42
	var b string = "Forty Two"

	/*reflect.TypeOf is the entry point of code introspection
	it receives and #EmptyInterface as an argument --check interfaces
	it returns a Type with a set of methods for code introspection
	*/

	typeA := reflect.TypeOf(a)
	fmt.Println(typeA)

	typeB := reflect.TypeOf(b)
	fmt.Println(typeB)

	fmt.Println()
	// for structs
	t := T{95, "Ninety Five"}

	typeT := reflect.TypeOf(t)
	fmt.Println(typeT)

	// Looping through the interface & accessing the Type and Value
	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		fmt.Println(field.Name, field.Type)
	}

	fmt.Println()
	/* checking if a type implements an Interface 
		using the Implements() method
	*/
	var pointerAdder *Adder
	adderType := reflect.TypeOf(pointerAdder).Elem()

	c := Calculator{}

	calculatorType := reflect.TypeOf(c)
	calculatorTypePointer := reflect.TypeOf(&c)

	// returns FALSE bcoz calculatorType does NOT implement the Adder Interface --its a pointer
	fmt.Println(calculatorType, calculatorType.Implements(adderType))

	// returns TRUE bcoz calculatorTypePointer does implement the Adder Interface --its a pointer
	fmt.Println(calculatorTypePointer, calculatorTypePointer.Implements(adderType))

}

// structure T
type T struct {
	A int32
	B string
}

// adder interface
type Adder interface {
	Add(int, int) int
}

type Calculator struct{}

// Method Add with a pointer receiver of type Calculator
func (c *Calculator) Add(a int, b int) int {
	return a + b
}
