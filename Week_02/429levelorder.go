package Week_02

type lNode struct {
	Val      int
	Children []*lNode
}

var lResult [][]int

func levelorder(root *lNode) [][]int {
	if root == nil {
		return nil
	}
	lResult = [][]int{}
	levelDfs(root, 0)
	return lResult
}

func levelDfs(root *lNode, level int) {
	if root == nil {
		return
	}

	if level == len(lResult) {
		lResult = append(lResult, []int{})
	}
	lResult[level] = append(lResult[level], root.Val)

	for _, n := range root.Children {
		levelDfs(n, level+1)
	}

}
