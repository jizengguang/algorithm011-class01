package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	dfs(root)
	return result
}

//递归
func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		result = append(result, root.Val)
		dfs(root.Right)
	}
}
