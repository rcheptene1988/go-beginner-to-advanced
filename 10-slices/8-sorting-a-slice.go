package main

import (
	"fmt"
	"sort"
)

// Sorting a slice

func main() {

	slice1 := []string{"Perl", "Java", "Go", "C++", "Ruby"}
	slice2 := []int{30, 20, 50, 90, 10, 40, 80, 60, 70}

	fmt.Println("Before sorting: ", slice1)

	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)

	sort.Strings(slice1)
	sort.Ints(slice2)

	fmt.Println("After sorting")

	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)
}
