package pkg

type TreeNode struct {
	Val    any
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(v any) *TreeNode {
	return &TreeNode{
		Val:    v,
		Height: 0,
		Left:   nil,
		Right:  nil,
	}
}

func sliceToTreeDFS(arr []any, i int) *TreeNode {
	if i < 0 || i >= len(arr) || arr[i] == nil {
		return nil
	}
	root := NewTreeNode(arr[i])
	root.Left = sliceToTreeDFS(arr, 2*i+1)
	root.Right = sliceToTreeDFS(arr, 2*i+2)
	return root
}

func SliceToTree(arr []any) *TreeNode {
	return sliceToTreeDFS(arr, 0)
}

func treetoSliceDFS(root *TreeNode, i int, res *[]any) {
	if root == nil {
		return
	}
	for i >= len(*res) {
		*res = append(*res, nil)
	}
	(*res)[i] = root.Val
	treetoSliceDFS(root.Left, 2*i+1, res)
	treetoSliceDFS(root.Right, 2*i+2, res)
}

func TreeToSlice(root *TreeNode) []any {
	var res []any
	treetoSliceDFS(root, 0, &res)
	return res
}
