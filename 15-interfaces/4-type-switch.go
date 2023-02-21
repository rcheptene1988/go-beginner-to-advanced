package main

import "fmt"

func main() {
	display(123)
	display("Go Programming Language")
	display(11.85)
	display(true)
}
func display(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("Type: %T -- Value: %v\n", a, a.(int))
	case string:
		fmt.Printf("Type: %T -- Value: %v\n", a, a.(string))
	case float64:
		fmt.Printf("Type: %T -- Value: %v\n", a, a.(float64))
	default:
		fmt.Println("Type not found")
	}

}
