package tree

import "letsgo/pkg"

type binarySearchTree struct {
	root *pkg.TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

func (bst *binarySearchTree) getRoot() *pkg.TreeNode {
	return bst.root
}

func (bst *binarySearchTree) insert(val int) {
	bst.root = bst.insertRec(bst.root, val)
}

func (bst *binarySearchTree) insertRec(node *pkg.TreeNode, val int) *pkg.TreeNode {
	if node == nil {
		return &pkg.TreeNode{Val: val}
	}
	if val < node.Val.(int) {
		node.Left = bst.insertRec(node.Left, val)
	} else {
		node.Right = bst.insertRec(node.Right, val)
	}
	return node
}

func (bst *binarySearchTree) print() {
	pkg.PrintTree(bst.root)
}
