package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := newBinarySearchTree()
	nums := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, num := range nums {
		bst.Insert(num)
	}

	bst.print()

	// bst.Remove(1)
	// bst.print()
	// bst.Remove(2)
	// bst.print()
	res := bst.inOrder()
	fmt.Println(res)
}
