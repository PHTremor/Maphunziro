package main

import "fmt"

func main() {
	// V1 ==================
	myDog := Dog{}
	myCat := Cat{}

	// RoarAndRun(&myDog) will not work -- coz we're passing a value not a pointer
	RoarAndRun(myDog)

	// RoarAndRun(myCat) will not work...why?
	RoarAndRun(&myCat)

	// V2 ==========================
	fmt.Println("\nEmpty interfaces ==========================")
	fmt.Println(aux)

	aux = 10
	fmt.Println(aux)

	aux = "Hello World"
	fmt.Println(aux)

	fmt.Println()
	// to know the type during runtime, use a switch statement
	emptyInterface := []interface{}{42, "Hello", true}
	for _, i := range emptyInterface {
		switch t := i.(type) {
		case int:
			fmt.Printf("%T ==> %d\n", t, i)
		case string:
			fmt.Printf("%T ==> %s\n", t, i)
		case bool:
			fmt.Printf("%T ==> %v\n", t, i)
		default:
			fmt.Printf("%T ==> %s\n", t, i)
		}
	}
}

// V1 of interfaces ======================================================================

/* Intefaces are methods or a collection of methods that define no logic or values
they defined methods to be implemented
*/

// #Choose one and not mixing the use of pointers and values -- choose consistency

// Animal interface has 2 methods
type Animal interface {
	Roar() string
	Run() string
}

// The Dog type defines the 2 Animal *methods
type Dog struct{}

// We pass a copy (value) of the Dog type
func (d Dog) Roar() string {
	return "woof!"
}

func (d Dog) Run() string {
	return "run like a dog"
}

// The Cat type defines the 2 Animal *methods
type Cat struct{}

// We pass a Pointer of the Cat type
func (c *Cat) Roar() string {
	return "meow!"
}

func (c *Cat) Run() string {
	return "run like a cat"
}

// This function received the Animal type to invoke the Roar and Run *methods
func RoarAndRun(a Animal) {
	fmt.Printf("%s and %s\n", a.Roar(), a.Run())
}

// V2 of interfaces ======================================================================
/* Empty interfaces -- it has no methods and is implemented by any type
useful when you cant know the data type beforehand
*/
var aux interface{}
