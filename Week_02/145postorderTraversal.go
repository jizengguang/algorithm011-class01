package Week_02

type PostTreeNode struct {
	Val   int
	Left  *PostTreeNode
	Right *PostTreeNode
}

var postResult []int

func postorderTraversal(root *PostTreeNode) []int {
	postResult = make([]int, 0)
	dfsPost(root)
	return postResult
}

func dfsPost(root *PostTreeNode) {
	if root != nil {
		dfsPost(root.Left)
		dfsPost(root.Right)
		postResult = append(postResult, root.Val)
	}

}
