package main

import "fmt"

// copy an array by value and reference into another array

func main() {
	array1 := [3]string{"kim", "jim", "bill"}

	array2 := array1

	array3 := &array1

	fmt.Printf("Array1: %v\n", array1)
	fmt.Printf("Array2: %v\n", array2)

	array1[0] = "Alex"

	fmt.Printf("Array1: %v\n", array1)
	fmt.Printf("Array2: %v\n", array2)

	fmt.Printf("Array3: %v\n", array3)

}
