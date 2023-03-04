package main

import "fmt"

func main() {
	var finger int = 1

	switch finger {
	case 0:
		fmt.Println("Thumb")
	case 1:
		fmt.Println("Index")
	case 2:
		fmt.Println("Middle")
	case 3:
		fmt.Println("Ring")
	case 4:
		fmt.Println("Pinkie")

	default:
		fmt.Println("Humans usually have five fingers")
	}
}

// ## takes a number && Prints name of the corresponding finger

//  You can also use condition on the cases like:
// 		case x < 2:
// 			fmt.Println("someLikeThat")

/*
You can stack cases when they have the same logic
	case 0:
	case 1:
		statement for 0 && 1
	default:
		default statment
*/

