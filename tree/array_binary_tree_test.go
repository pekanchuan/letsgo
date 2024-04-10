package tree

import (
	"fmt"
	"letsgo/pkg"
	"testing"
)

func TestArrayBinaryTree(t *testing.T) {
	arr := []any{1, 2, 3, 4, nil, 6, 7, 8, 9, nil, nil, 12, nil, nil, 15}
	root := pkg.SliceToTree(arr)
	fmt.Println(arr...)
	pkg.PrintTree(root)

	abt := newArrayBinaryTree(arr)

	i := 1
	l := abt.left(i)
	r := abt.right(i)
	p := abt.parent(i)
	fmt.Println("\n当前节点的索引为", i, "，值为", abt.val(i))
	fmt.Println("其左子节点的索引为", l, "，值为", abt.val(l))
	fmt.Println("其右子节点的索引为", r, "，值为", abt.val(r))
	fmt.Println("其父节点的索引为", p, "，值为", abt.val(p))

	res := abt.levelOrder()
	fmt.Println("level order:", res)
	res = abt.preOrder()
	fmt.Println("pre order", res)
}
