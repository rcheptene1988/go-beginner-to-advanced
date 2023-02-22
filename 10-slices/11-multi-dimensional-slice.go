package main

import "fmt"

// Multi-dimensional slices  - are exactly the same as arrays, just the slices don't contain the size

func main() {

	slice1 := [][]int{
		{10, 20},
		{30, 40, 50},
		{60, 70, 80, 90},
		{100, 120},
	}

	fmt.Println(slice1)
}
