package main

import (
	"datastructures/stack"
	"fmt"
)

func main() {
	var stack structure.Stack[int]

	stack.Add(10)
	stack.Add(20)
	stack.Add(30)

	for !stack.IsNullOrEmpty() {
		value, err := stack.Pop()

		if err != nil {
			fmt.Println(err)
		}

		println(value)
	}
}
