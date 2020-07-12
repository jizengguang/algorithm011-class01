package Week_03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//前序的第一个，肯定是根节点
	//中序的前面，肯定是左子树，然后是根节点。
	//根据前序的根节点，找到中序根节点的位置，把左子树的构建出来。

	root := &TreeNode{preorder[0], nil, nil}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	//找到中序遍历的根结点i，从0到i-1，即左子树。

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root

}
