package main

import "fmt"

// fixed-length size
// var array_name[length] type
func main() {
	var names [3]string

	names[0] = "Roman"
	names[1] = "Claudia"
	names[2] = "Paul"

	fmt.Println(names)

}
