package Week_04

//Definition for a binary tree node.
type LargeTreeNode struct {
	Val   int
	Left  *LargeTreeNode
	Right *LargeTreeNode
}

func largestValues(root *LargeTreeNode) []int {
	if root == nil {
		return nil
	}

	var queue = make([]*LargeTreeNode, 0)
	queue = append(queue, root)
	var res = make([]int, 0)

	for len(queue) > 0 {
		var length = len(queue)
		var max = -1 << 63

		for i := 0; i < length; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, max)
		queue = queue[length:]

	}
	return res

}
