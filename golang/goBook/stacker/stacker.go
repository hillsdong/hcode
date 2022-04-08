package main

import (
	"fmt"
	"learn/gobook/src/stacker/stack"
)

func main() {
	var haystack stack.Stack

	haystack.Push(1.0)
	haystack.Push("hills")
	haystack.Push([]string{"a", "b", "c"})

	for {
		v, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(v)
	}
}
