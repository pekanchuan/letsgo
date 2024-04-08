package pkg

import (
	"container/list"
	"fmt"
)

func PrintSlice[T any](nums []T) {
	fmt.Printf("%v", nums)
	fmt.Println()
}

func PrintList(list *list.List) {
	if list.Len() == 0 {
		fmt.Print("[]\n")
	}
	e := list.Front()
	fmt.Print("[")
	for e.Next() != nil {
		fmt.Print(e.Value, " ")
		e = e.Next()
	}
	fmt.Print(e.Value, "]\n")
}

func PrintMap[K comparable, V any](m map[K]V) {
	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}

func PrintHeap(h []any) {
	fmt.Printf("堆的数组表示: ")
	fmt.Printf("%v", h)
	fmt.Printf("\n堆的树状表示: ")
	root := SliceToTree(h)
	PrintTree(root)
}

func PrintLinkedList() {}

func PrintTree(root *TreeNode) {
	printTreeHelper(root, nil, false)
}

func printTreeHelper(root *TreeNode, prev *trunk, isRight bool) {
	if root == nil {
		return
	}
	prevStr := "    "
	trunk := newTrunk(prev, prevStr)
	printTreeHelper(root.Right, trunk, true)
	if prev == nil {
		trunk.str = "———"
	} else if isRight {
		trunk.str = "/———"
		prevStr = "   |"
	} else {
		trunk.str = "\\———"
		prev.str = prevStr
	}
	showTrunk(trunk)
	fmt.Println(root.Val)
	if prev != nil {
		prev.str = prevStr
	}
	trunk.str = "   |"
	printTreeHelper(root.Left, trunk, false)
}

type trunk struct {
	prev *trunk
	str  string
}

func newTrunk(prev *trunk, str string) *trunk {
	return &trunk{
		prev: prev,
		str:  str,
	}
}

func showTrunk(t *trunk) {
	if t == nil {
		return
	}
	showTrunk(t.prev)
	fmt.Print(t.str)
}
