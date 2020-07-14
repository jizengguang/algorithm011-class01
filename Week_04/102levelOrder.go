package Week_04

type levelOrderTreeNode struct {
	Val   int
	Left  *levelOrderTreeNode
	Right *levelOrderTreeNode
}

/*
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
*/

//刻意练习dfs
func levelOrder(root *levelOrderTreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue = make([]*levelOrderTreeNode, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		ans := make([]int, 0)
		length := len(queue)

		for i := 0; i < length; i++ {
			node := queue[i]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, ans)
		queue = queue[length:]

	}
	return res

}
