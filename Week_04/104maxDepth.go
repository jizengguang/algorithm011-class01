package Week_04

import "math"

type maxDepthTreeNode struct {
	Val   int
	Left  *maxDepthTreeNode
	Right *maxDepthTreeNode
}

func maxDepth(root *maxDepthTreeNode) int {
	if root == nil {
		return 0
	}
	//加1 是为了加上根节点那一层
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
}
