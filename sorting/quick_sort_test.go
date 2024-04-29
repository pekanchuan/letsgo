package sorting

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	res := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	quickSort(res, 0, 9)

	fmt.Println(res)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
