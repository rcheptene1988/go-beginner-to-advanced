// Access the underlying variable of interface
// can be accessed in 2 ways: type assertion and type switch

// type assertion syntax:
// * a.(T)   // a - the value of the interface, and T - is the type, known as the asserted type
package main

import "fmt"

func main() {
	var result = "Go Programming Language"

	display(result)
}

func display(a interface{}) {
	value := a.(string)
	fmt.Println("value: ", value)
}
