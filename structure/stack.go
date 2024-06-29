package structure

import "errors"

type stackNode[T any] struct {
	PreviousNode *stackNode[T]
	Value        T
}

type Stack[T any] struct {
	topNode *stackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		topNode: nil,
	}
}

func (s *Stack[T]) Add(newValue T) *Stack[T] {
	s.topNode = &stackNode[T]{
		PreviousNode: s.topNode,
		Value:        newValue,
	}

	return s
}

func (s *Stack[T]) Pop() (T, error) {
	if s.topNode == nil {
		var zero T
		return zero, errors.New("cannot pop from an empty structure")
	}

	value := s.topNode.Value
	s.topNode = s.topNode.PreviousNode

	return value, nil
}

func (s *Stack[T]) IsNullOrEmpty() bool {
	return (s == nil) || (s.topNode == nil)
}
