package structure

import (
	"errors"
	"math/bits"
)

type MinMaxHeap[T any] struct {
	elements []T
	compare  func(T, T) int
	length   int
}

func NewMinMaxHeap[T any](initialElements []T, compareFunc func(T, T) int) *MinMaxHeap[T] {
	heap := MinMaxHeap[T]{
		elements: make([]T, len(initialElements)),
		compare:  compareFunc,
		length:   len(initialElements),
	}
	copy(heap.elements, initialElements)

	heap.restoreHeapProperties()

	return &heap
}

func level(idx int) int {
	return bits.Len(uint(idx+1)) - 1
}

func parent(idx int) int {
	if idx == 0 {
		return 0
	}
	return (idx - 1) >> 1
}

func (minMaxHeap *MinMaxHeap[T]) swap(i int, j int) {
	minMaxHeap.elements[i], minMaxHeap.elements[j] = minMaxHeap.elements[j], minMaxHeap.elements[i]
}

func (minMaxHeap *MinMaxHeap[T]) isValidIdx(idx int) bool {
	return idx < minMaxHeap.length
}

func (minMaxHeap *MinMaxHeap[T]) subtree(idx int, depth int) []int {
	if !minMaxHeap.isValidIdx(idx) {
		return []int{}
	}

	subtree := []int{idx}

	if depth != 0 {
		subtree = append(subtree, minMaxHeap.subtree((idx<<1)+1, depth-1)...)
		subtree = append(subtree, minMaxHeap.subtree((idx<<1)+2, depth-1)...)
	}

	return subtree
}

func (minMaxHeap *MinMaxHeap[T]) getMin(indexes []int) int {
	if len(indexes) == 0 {
		return -1
	}

	minIdx := indexes[0]
	for _, idx := range indexes[1:] {
		if minMaxHeap.compare(minMaxHeap.elements[minIdx], minMaxHeap.elements[idx]) > 0 {
			minIdx = idx
		}
	}

	return minIdx
}

func (minMaxHeap *MinMaxHeap[T]) getMax(indexes []int) int {
	if len(indexes) == 0 {
		return -1
	}

	maxIdx := indexes[0]
	for _, idx := range indexes[1:] {
		if minMaxHeap.compare(minMaxHeap.elements[maxIdx], minMaxHeap.elements[idx]) < 0 {
			maxIdx = idx
		}
	}

	return maxIdx
}

func (minMaxHeap *MinMaxHeap[T]) restoreHeapProperties() {
	for i := (minMaxHeap.length >> 1) - 1; i >= 0; i-- {
		if level(i)%2 == 0 {
			minMaxHeap.pushDownMin(i)
		} else {
			minMaxHeap.pushDownMax(i)
		}
	}
}

func (minMaxHeap *MinMaxHeap[T]) pushDownMin(idx int) {
	if !minMaxHeap.isValidIdx(idx) {
		return
	}

	subtree := minMaxHeap.subtree(idx, 2)
	minIdx := minMaxHeap.getMin(subtree)

	if minIdx == idx {
		return
	} else if level(minIdx)-level(idx) == 2 {
		minMaxHeap.swap(idx, minIdx)

		if par := parent(minIdx); minMaxHeap.compare(minMaxHeap.elements[par], minMaxHeap.elements[minIdx]) < 0 {
			minMaxHeap.swap(par, minIdx)
		}
	} else {
		minMaxHeap.swap(idx, minIdx)
	}
}

func (minMaxHeap *MinMaxHeap[T]) pushDownMax(idx int) {
	if !minMaxHeap.isValidIdx(idx) {
		return
	}

	subtree := minMaxHeap.subtree(idx, 2)
	maxIdx := minMaxHeap.getMax(subtree)

	if maxIdx == idx {
		return
	} else if level(maxIdx)-level(idx) == 2 {
		minMaxHeap.swap(idx, maxIdx)

		if par := parent(maxIdx); minMaxHeap.compare(minMaxHeap.elements[maxIdx], minMaxHeap.elements[par]) < 0 {
			minMaxHeap.swap(par, maxIdx)
		}
	} else {
		minMaxHeap.swap(idx, maxIdx)
	}
}

func (minMaxHeap *MinMaxHeap[T]) pushUp(idx int) {
	if idx == 0 {
		return
	}

	par := parent(idx)

	if level(idx)%2 == 0 {
		if minMaxHeap.compare(minMaxHeap.elements[par], minMaxHeap.elements[idx]) < 0 {
			minMaxHeap.swap(idx, par)
			minMaxHeap.pushUpMax(par)
		} else {
			minMaxHeap.pushUpMin(idx)
		}
	} else {
		if minMaxHeap.compare(minMaxHeap.elements[idx], minMaxHeap.elements[par]) < 0 {
			minMaxHeap.swap(idx, par)
			minMaxHeap.pushUpMin(par)
		} else {
			minMaxHeap.pushUpMax(idx)
		}
	}
}

func (minMaxHeap *MinMaxHeap[T]) pushUpMin(idx int) {
	grandParent := parent(parent(idx))

	if level(idx)-level(grandParent) != 2 {
		return
	}

	if minMaxHeap.compare(minMaxHeap.elements[idx], minMaxHeap.elements[grandParent]) < 0 {
		minMaxHeap.swap(idx, grandParent)
		minMaxHeap.pushUpMin(grandParent)
	}
}

func (minMaxHeap *MinMaxHeap[T]) pushUpMax(idx int) {
	grandParent := parent(parent(idx))

	if level(idx)-level(grandParent) != 2 {
		return
	}

	if minMaxHeap.compare(minMaxHeap.elements[grandParent], minMaxHeap.elements[idx]) < 0 {
		minMaxHeap.swap(idx, grandParent)
		minMaxHeap.pushUpMax(grandParent)
	}
}

func (minMaxHeap *MinMaxHeap[T]) Insert(element T) {
	minMaxHeap.elements = append(minMaxHeap.elements, element)
	minMaxHeap.length++

	minMaxHeap.pushUp(minMaxHeap.length - 1)
}

func (minMaxHeap *MinMaxHeap[T]) Min() (T, error) {
	if minMaxHeap.length == 0 {
		var zero T
		return zero, errors.New("cannot get min of an empty heap")
	}

	minElement := minMaxHeap.elements[0]

	minMaxHeap.elements[0] = minMaxHeap.elements[minMaxHeap.length-1]
	minMaxHeap.elements = minMaxHeap.elements[:minMaxHeap.length-1]
	minMaxHeap.length--

	minMaxHeap.pushDownMin(0)

	return minElement, nil
}

func (minMaxHeap *MinMaxHeap[T]) Max() (T, error) {
	if minMaxHeap.length == 0 {
		var zero T
		return zero, errors.New("cannot get min of an empty heap")
	}

	if minMaxHeap.length == 1 {
		return minMaxHeap.Min()
	}

	maxElementIdx := 1
	if (minMaxHeap.length > 2) && (minMaxHeap.compare(minMaxHeap.elements[1], minMaxHeap.elements[2]) < 0) {
		maxElementIdx = 2
	}

	maxElement := minMaxHeap.elements[maxElementIdx]

	minMaxHeap.elements[maxElementIdx] = minMaxHeap.elements[minMaxHeap.length-1]
	minMaxHeap.elements = minMaxHeap.elements[:minMaxHeap.length-1]
	minMaxHeap.length--

	for i := maxElementIdx; i >= 0; i-- {
		if level(i)%2 == 0 {
			minMaxHeap.pushDownMin(i)
		} else {
			minMaxHeap.pushDownMax(i)
		}
	}

	return maxElement, nil
}
