package tree

import (
	"fmt"
	"letsgo/pkg"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := pkg.SliceToTree([]any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("初始化二叉树:")
	pkg.PrintTree(root)

	nums := levelOrder(root)
	fmt.Println("层序遍历的节点打印序列 =", nums)
}
