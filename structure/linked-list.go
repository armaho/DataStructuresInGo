package structure

import "errors"

type LinkedListNode[T comparable] struct {
	PreviousNode *LinkedListNode[T]
	NextNode     *LinkedListNode[T]
	Value        T
}

type LinkedList[T comparable] struct {
	SentinelNode *LinkedListNode[T]
	length       int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	var zero T

	sentinelNode := &LinkedListNode[T]{
		PreviousNode: nil,
		NextNode:     nil,
		Value:        zero,
	}

	sentinelNode.NextNode = sentinelNode
	sentinelNode.PreviousNode = sentinelNode

	return &LinkedList[T]{
		SentinelNode: sentinelNode,
		length:       0,
	}
}

func (linkedList *LinkedList[T]) Insert(value T, node *LinkedListNode[T]) *LinkedListNode[T] {
	newNode := &LinkedListNode[T]{
		PreviousNode: node,
		NextNode:     node.NextNode,
		Value:        value,
	}

	node.NextNode = newNode
	newNode.NextNode.PreviousNode = newNode

	linkedList.length++

	return newNode
}

func (linkedList *LinkedList[T]) Delete(node *LinkedListNode[T]) error {
	if node == linkedList.SentinelNode {
		return errors.New("cannot delete the sentinel node from a linked list")
	}

	node.PreviousNode.NextNode = node.NextNode
	node.NextNode.PreviousNode = node.PreviousNode

	linkedList.length--

	return nil
}

func (linkedList *LinkedList[T]) Search(value T) *LinkedListNode[T] {
	if linkedList == nil {
		return nil
	}

	for node := linkedList.SentinelNode.NextNode; node != linkedList.SentinelNode; node = node.NextNode {
		if node.Value == value {
			return node
		}
	}

	return nil
}

func (linkedList *LinkedList[T]) Length() int {
	return linkedList.length
}
