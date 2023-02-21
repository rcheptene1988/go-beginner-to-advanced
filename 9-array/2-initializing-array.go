package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3} // initializing and array with ellipses ... - We can use them instead of specifying the length. The compiler can identify the length of an array based on
	// the elements specified in the array
	fmt.Println("Length of Array: ", len(x))

	y := [5]int{1: 10, 3: 30} // assign valut to specific key
	fmt.Println(y)

	// filter array
	z := [5]string{"Kim", "Jim", "Bob", "Robert", "David"}
	fmt.Println("Name: %v\n", z)
	fmt.Println(z[:])
	fmt.Println(z[:3])
	fmt.Println(z[2:4])

}
