package main

import "fmt"

// Iterate over a slice
// There are different ways
// Using for loop
// Using range in for loop
// Using a blank identifier
// Using for loop as a while

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70}

	// For loop
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d\t", slice[i])
	}
	fmt.Println("")

	// Use range in for loop
	for index, value := range slice {
		fmt.Printf("Index = %d and Value = %d\n", index, value)
	}

	// Using range in for loop, without the index
	for _, value := range slice {
		fmt.Printf("Value = %d\n", value)
	}

	// Using for loop as while
	i := 0
	for range slice {
		fmt.Println(slice[i])
		i++
	}
}
