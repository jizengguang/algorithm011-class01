package Week_09

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	if left == -1 {
		return -1
	}

	right := height(root.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(left-right) > 1 {
		return -1
	}
	return math.Max(left, right) + 1

}
