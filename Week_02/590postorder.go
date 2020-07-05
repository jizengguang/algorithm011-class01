package Week_02

type pNode struct {
	Val      int
	Children []*pNode
}

var postOrderResult []int

func postorder(root *pNode) []int {
	postOrderResult = make([]int, 0)
	postOrderDfs(root)
	return postOrderResult
}

func postOrderDfs(root *pNode) {
	if root != nil {
		for _, n := range root.Children {
			postOrderDfs(n)
		}
		postOrderResult = append(postOrderResult, root.Val)
	}
}
