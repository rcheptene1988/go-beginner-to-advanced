package main

import "fmt"

// Add items to existing slice using append function

func main() {

	slice := make([]int, 2, 5)
	slice[0] = 10
	slice[1] = 20

	fmt.Println("Slice1: ", slice)

	fmt.Printf("Length: %d\tCapacity: %d\t\n", len(slice), cap(slice))

	slice = append(slice, 30, 40, 50, 60, 70, 80, 90, 100, 110)

	fmt.Println("Slice 1 after appending data: ", slice)

	fmt.Printf("Length: %d\tCapacity: %d\t\n", len(slice), cap(slice))
}
