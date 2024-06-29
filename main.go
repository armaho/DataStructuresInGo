package main

import (
	"datastructures/structure"
	"fmt"
)

func main() {
	linkedList := structure.NewLinkedList[int]()

	firstNode := linkedList.Insert(10, linkedList.SentinelNode)
	thirdNode := linkedList.Insert(20, firstNode)
	linkedList.Insert(30, firstNode)

	fmt.Println(linkedList.Search(20))
	fmt.Println(thirdNode)

	err := linkedList.Delete(thirdNode)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(linkedList.Search(20))
}
