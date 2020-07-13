package Week_04

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	//1、迭代
	//时间复杂度O(N)
	//空间复杂度O(N)
	return dfs(root, nil, nil)
}

func dfs(root, low, upper *TreeNode) bool {
	if root == nil {
		return true
	}
	if low != nil && root.Val <= low.Val || upper != nil && root.Val >= upper.Val {
		return false
	}
	return dfs(root.Left, low, root) && dfs(root.Right, root, upper)
}

//2、中序遍历是递增的

