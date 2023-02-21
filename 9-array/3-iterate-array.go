package main

import "fmt"

//iterate over an array

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}
	// range
	for index, value := range intArray {
		fmt.Println(index, "=>", value)
	}

	//range but no index
	for _, value := range intArray {
		fmt.Println(value)
	}

	// range as a while loop
	j := 0
	for range intArray {
		fmt.Println("While loop", intArray[j])
		j++
	}

}
