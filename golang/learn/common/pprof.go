package main

import "fmt"

func main() {
	var a int = 1
	do(a)
}
func do(a interface{}) {
	//	a = nil

	switch a.(type) {
	case int, interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("null")

	}
	if a == nil {
		fmt.Println("nil")

	}
	fmt.Println(a)
}
