package tree

import (
	"fmt"
	"letsgo/pkg"
	"reflect"
	"testing"
)

func TestPreOrder(t *testing.T) {
	root := pkg.SliceToTree([]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("初始化二叉树:")
	pkg.PrintTree(root)

	nums = nil
	preOrder(root)
	fmt.Println("前序遍历的节点打印序列 =", nums)
	expectation := []any{0, 1, 3, 7, 8, 4, 9, 2, 5, 6}
	if !reflect.DeepEqual(nums, expectation) {
		t.Errorf("预期的前序遍历序列是 %v, 但得到的是 %v", expectation, nums)
	}
}

func TestInOrder(t *testing.T) {
	root := pkg.SliceToTree([]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("初始化二叉树:")
	pkg.PrintTree(root)

	nums = nil
	inOrder(root)
	fmt.Println("中序遍历的节点打印序列 =", nums)
	expectation := []any{7, 3, 8, 1, 9, 4, 0, 5, 2, 6}
	if !reflect.DeepEqual(nums, expectation) {
		t.Errorf("预期的中序遍历序列是 %v, 但得到的是 %v", expectation, nums)
	}
}

func TestPostOrder(t *testing.T) {
	root := pkg.SliceToTree([]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("初始化二叉树:")
	pkg.PrintTree(root)

	nums = nil
	postOrder(root)
	fmt.Println("后序遍历的节点打印序列 =", nums)
	expectation := []any{7, 8, 3, 9, 4, 1, 5, 6, 2, 0}
	if !reflect.DeepEqual(nums, expectation) {
		t.Errorf("预期的后序遍历序列是 %v, 但得到的是 %v", expectation, nums)
	}
}
