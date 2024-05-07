package heap

import (
	"fmt"
	. "letsgo/pkg"
	"testing"
)

type maxHeap struct {
	data []any
}

func newHeap() *maxHeap {
	return &maxHeap{
		data: make([]any, 0),
	}
}

func newMaxHeap(nums []any) *maxHeap {
	h := &maxHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *maxHeap) left(i int) int {
	return 2*i + 1
}

func (h *maxHeap) right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *maxHeap) size() int {
	return len(h.data)
}

func (h *maxHeap) isEmpty() bool {
	return len(h.data) == 0
}

func (h *maxHeap) peek() any {
	return h.data[0]
}

func (h *maxHeap) push(val int) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *maxHeap) siftUp(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.data[i].(int) <= h.data[p].(int) {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *maxHeap) pop() any {
	if h.isEmpty() {
		fmt.Println("error")
		return nil
	}
	h.swap(0, h.size()-1)
	val := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)
	return val
}

func (h *maxHeap) siftDown(i int) {
	for {
		l, r, max := h.left(i), h.right(i), i
		if l < h.size() && h.data[l].(int) > h.data[max].(int) {
			max = l
		}
		if r < h.size() && h.data[r].(int) > h.data[max].(int) {
			max = r
		}
		if max == i {
			break
		}
		h.swap(i, max)
		i = max
	}
}

func (h *maxHeap) print() {
	PrintHeap(h.data)
}

func TestMyHeap(t *testing.T) {
	maxHeap := newMaxHeap([]any{9, 8, 6, 6, 7, 5, 2, 1, 4, 3, 6, 2})
	fmt.Println("After inputting an array and building a heap")
	maxHeap.print()

	peek := maxHeap.peek()
	fmt.Println("top element is ", peek)

	val := 7
	maxHeap.push(val)
	maxHeap.print()

	peek = maxHeap.pop()
	maxHeap.print()

}
