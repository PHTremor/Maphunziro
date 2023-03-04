package main

import "fmt"

func main() {
	// Function 1 ==============================================================================
	fmt.Println("Function 1 =============================================")

	sum, sub := ops(2, 2)
	fmt.Println("2+2=", sum, "2-2=", sub)

	b, _ := ops(10, 2)
	fmt.Println("10+2=", b)

	// Function 2 ==============================================================================
	fmt.Println("Function 2 =============================================")

	total := sumFunc(1, 2, 3, 4, 5)
	fmt.Println("The sum of the the first 5 numbers is", total)

	// Function 3 ==============================================================================
	fmt.Println("Function 3 =============================================")

	c := doit(addition, 2, 3)
	fmt.Println("2+3=", c)

	d := doit(multiply, 2, 3)
	fmt.Println("2*3=", d)

	// Function 4 ==============================================================================
	fmt.Println("Function 4 =============================================")

	e := accumulator(1)
	f := accumulator(2)

	fmt.Println("e", "f")
	for i:=0; i<5; i++ {
		fmt.Println(e(), f())
	}

}

// Function 1
func ops(a int, b int) (int, int) {
	return a + b, a - b
}

// Function 2 => Variadic
// Recieves undetermined number of arguments... computes their sum
func sumFunc(nums ...int) int {
	total := 0
	for _, a := range nums {
		total = total + a
	}
	return total
}

// Function 3 => Functions as arguments
// Receives 3 args: a func, & 2 int values
func doit(operator func(int, int) int, a int, b int) int {
	return operator(a, b)
}

func addition(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

// Function 4 => anonymous functions AKA a closure
// A closure can refer to variables outside its body
func accumulator(increment int) func() int  {
	i := 0

	return func() int {
		i = i + increment
		
		return i
	}
}