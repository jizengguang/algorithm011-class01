package Week_04

type geneTreeNode struct {
	Val   int
	Left  *geneTreeNode
	Right *geneTreeNode
}

func generateTrees(n int) []*geneTreeNode {
	if n == 0 {
		return nil
	}
	return gene(1, n)
}

func gene(start, end int) []*geneTreeNode {
	if start > end {
		return []*geneTreeNode{nil}
	}
	var newTree = make([]*geneTreeNode, 0)

	for i := start; i < end; i++ {

		//获取左右子树
		leftTrees := gene(start, i-1)
		rightTrees := gene(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				//设置根节点
				var tmp = &geneTreeNode{i, nil, nil}
				tmp.Left = left
				tmp.Right = right
				newTree = append(newTree, tmp)
			}
		}

	}
	return newTree
}
