package tree

import . "letsgo/pkg"

type aVLTree struct {
	root *TreeNode
}

func newAVLTree() *aVLTree {
	return &aVLTree{
		root: nil,
	}
}

func (t *aVLTree) getHeight(node *TreeNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}
