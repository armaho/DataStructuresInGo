package main

import (
	"datastructures/structure"
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	heap := structure.NewMinMaxHeap(s1, func(a int, b int) int { return a - b })

	heap.Insert(10)
	heap.Insert(0)

	fmt.Println(heap.Max())
	fmt.Println(heap.Min())
	fmt.Println(heap.Max())
	fmt.Println(heap.Min())
}
