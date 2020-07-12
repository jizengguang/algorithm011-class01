package Week_03

//Definition for TreeNode.
type LowTreeNode struct {
	Val   int
	Left  *LowTreeNode
	Right *LowTreeNode
}

func lowestCommonAncestor(root, p, q *LowTreeNode) *LowTreeNode {
	//terminator
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//在左子树中存在 且在右子树中存在，那么当前就是公共祖先，return返回。
	if left != nil && right != nil {
		return root
	}

	//如果在左子树中不存在，那么一定在右子树中。
	if left == nil {
		return right
	}

	return right

}
