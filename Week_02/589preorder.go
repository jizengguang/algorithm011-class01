package Week_02

type Node struct {
	Val      int
	Children []*Node
}

var preOrderResult []int

func preorder(root *Node) []int {
	preOrderResult = make([]int, 0)
	preOrderDfs(root)
	return preOrderResult
}

func preOrderDfs(root *Node) {
	if root != nil {
		preOrderResult = append(preOrderResult, root.Val)
		for _, value := range root.Children {
			preOrderDfs(value)
		}
	}
}
