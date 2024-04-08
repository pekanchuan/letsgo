package tree

import "letsgo/pkg"

var nums []any

func preOrder(node *pkg.TreeNode) {
	if node == nil {
		return
	}
	nums = append(nums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

func inOrder(node *pkg.TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	nums = append(nums, node.Val)
	inOrder(node.Right)
}

func postOrder(node *pkg.TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	nums = append(nums, node.Val)
}
