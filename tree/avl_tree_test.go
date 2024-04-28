package tree

import (
	"fmt"
	. "letsgo/pkg"
	"testing"
)

func TestAVLTree(t *testing.T) {
	tree := newAVLTree()
	testInsert(tree, 1)
	testInsert(tree, 2)
	testInsert(tree, 3)
	testInsert(tree, 4)
	testInsert(tree, 5)
	testInsert(tree, 8)
	testInsert(tree, 7)
	testInsert(tree, 9)
	testInsert(tree, 10)
	testInsert(tree, 6)

	testInsert(tree, 7)

	testRemove(tree, 8)
	testRemove(tree, 5)
	testRemove(tree, 4)

	node := tree.search(7)
	fmt.Printf("\n查找到的节点对象为 %#v , 节点值 = %d \n", node, node.Val)
}

func testInsert(tree *aVLTree, val int) {
	tree.insert(val)
	fmt.Printf("\n插入节点 %d 后, AVL 树为 \n", val)
	PrintTree(tree.root)
}

func testRemove(tree *aVLTree, val int) {
	tree.remove(val)
	fmt.Printf("\n删除节点 %d 后, AVL 树为 \n", val)
	PrintTree(tree.root)
}
