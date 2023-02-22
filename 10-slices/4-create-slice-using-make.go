package main

import "fmt"

// Create slice unsing make function
// Sintax: func make([]Type, len, cap) []Type    - cap (capacity) is optional
//* Generally, make function is used to create empty slices

func main() {
	var slice1 = make([]int, 4, 7)

	fmt.Printf("Slice1 = %v, \tLength = %d, \tCapacity = %d\n", slice1, len(slice1), cap(slice1))
}
