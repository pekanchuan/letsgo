package tree

import (
	. "letsgo/pkg"
)

type binarySearchTree struct {
	root *TreeNode
}

func newBinarySearchTree() *binarySearchTree {
	bst := &binarySearchTree{}
	bst.root = nil
	return bst
}

func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

func (bst *binarySearchTree) search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if num < node.Val.(int) {
			node = node.Left
		} else if num > node.Val.(int) {
			node = node.Right
		} else {
			return node
		}
	}
	return nil
}

func (bst *binarySearchTree) Insert(value int) {
	bst.root = bst.insert(bst.root, value)
}

func (bst *binarySearchTree) insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return NewTreeNode(value)
	}

	if value < node.Val.(int) {
		node.Left = bst.insert(node.Left, value)
	} else {
		node.Right = bst.insert(node.Right, value)
	}
	return node
}

func (bst *binarySearchTree) Remove(value int) {
	bst.root = bst.removeNode(bst.root, value)
}

func (bst *binarySearchTree) removeNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.Val.(int) {
		node.Left = bst.removeNode(node.Left, value)
	} else if value > node.Val.(int) {
		node.Right = bst.removeNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minNode := bst.findMin(node.Right)
		node.Val = minNode.Val
		node.Right = bst.removeNode(node.Right, minNode.Val.(int))
	}
	return node
}

func (bst *binarySearchTree) findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *binarySearchTree) inOrder() (res []int) {
	bst.inOrderHelper(bst.root, &res)
	return
}

func (bst *binarySearchTree) inOrderHelper(node *TreeNode, res *[]int) {
	if node != nil {
		bst.inOrderHelper(node.Left, res)
		*res = append(*res, node.Val.(int))
		bst.inOrderHelper(node.Right, res)
	}
}

func (bst *binarySearchTree) print() {
	PrintTree(bst.root)
}
