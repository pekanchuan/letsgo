package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type intHeap []any

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i].(int) > (*h)[j].(int)
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Top() any {
	return (*h)[0]
}

func TestHeap(t *testing.T) {

	maxHeap := &intHeap{}
	heap.Init(maxHeap)

	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	top := maxHeap.Top()
	fmt.Printf("top heap element is %d\n", top)

	heap.Pop(maxHeap)
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)
	heap.Pop(maxHeap)

	size := len(*maxHeap)
	fmt.Printf("heap elements count is %d\n", size)

	isEmpty := len(*maxHeap) == 0
	fmt.Printf("is heap empty %t\n", isEmpty)
}
