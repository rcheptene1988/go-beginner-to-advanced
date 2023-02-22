package main

import "fmt"

//  Append a slice to an existing slice

func main() {
	var slice1 = []string{"Java", "C#", "Go"}
	var slice2 = []string{"Perl", "PHP"}

	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	slice2 = append(slice2, slice1...)

	fmt.Println("After append")
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)
}
