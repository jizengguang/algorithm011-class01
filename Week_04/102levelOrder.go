package Week_04

type levelOrderTreeNode struct {
	Val   int
	Left  *levelOrderTreeNode
	Right *levelOrderTreeNode
}

var res [][]int

func levelOrder(root *levelOrderTreeNode) [][]int {
	if root == nil {
		return nil
	}

	res = make([][]int, 0)
	bfs(root, 0)
	return res
}

func bfs(root *levelOrderTreeNode, level int) {
	if root == nil {
		return
	}

	if level == len(res) {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)

	bfs(root.Left, level+1)
	bfs(root.Right, level+1)
}
