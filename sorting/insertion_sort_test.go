package sorting

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	insertionSort(nums)
	fmt.Println(nums)
}

func insertionSort(nums []int) {
	n := len(nums)

	for i := 1; i < n; i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}
