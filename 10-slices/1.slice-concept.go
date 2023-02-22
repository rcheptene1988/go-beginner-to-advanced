package main

import "fmt"

// Slices store only elements of similar type
// A slice is a reference to an underlying array
// Unlike arrays, the size of the slices is flexible and can be changed

// Slices are represented by 3 things
// 1 Pointers: pointer to the underlying array
// 2 Length: Current length of the underlying array
// 3 Capacity: Total capacity which is the maximum capacity to which the underlying array can expand

//* A slice is declared just like and array, but it doesn't contain the size of the slice, so it can grow or shrink according to the requirements

// Define slices
// Create slice using slice literal: var mySlice = []string{"a", "b", "c"} //* REMEMBER when you create a slice using a string literal, it first creates an array, and after that returns a slice reference to it.
// Create a slice

func main() {
	var slice1 = []string{"a", "b", "c"}
	fmt.Println(slice1)

	slice2 := []string{"e", "f", "g"}
	fmt.Println(slice2)

	for i := 0; i < len(slice2); i++ {
		fmt.Printf("Slice[%d] = %s\t", i, slice2[i])
	}
	fmt.Println()

}
