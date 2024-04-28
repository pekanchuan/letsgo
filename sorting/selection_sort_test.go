package sorting

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	foo := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := selectionSort(foo)

	fmt.Println(res)
}

func selectionSort(list []int) (res []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if list[j] < list[k] {
				k = j
			}
		}
		list[i], list[k] = list[k], list[i]
	}
	res = list
	return
}
