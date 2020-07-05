package Week_02

type PreTreeNode struct {
	Val   int
	Left  *PreTreeNode
	Right *PreTreeNode
}

var preResult []int

func preorderTraversal(root *PreTreeNode) []int {
	preResult = make([]int, 0)
	dfsorder(root)
	return preResult
}

func dfsorder(root *PreTreeNode) {
	if root != nil {
		preResult = append(preResult, root.Val)
		dfsorder(root.Left)
		dfsorder(root.Right)
	}
}
