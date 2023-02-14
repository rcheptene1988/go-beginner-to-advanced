package main

import "fmt"

func main() {
	var i int = 100
	var s string = "Hello"

	kind(i)
	kind(s)
}

func kind(i interface{}) {
	fmt.Printf("The kind if %v is %T\n", i, i)
}
