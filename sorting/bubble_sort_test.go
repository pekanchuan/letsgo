package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{9, 8, 7, 6, 4, 5, 3, 2, 1, 0}
	bubbleSort(nums)

	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}
