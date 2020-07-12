学习笔记

#作业
1、二叉树的公共祖先：  
定义f(x) 表示x节点的子树中是否包含p,q的节点。如果包含即为true，反之false。  
怎么判断f(x)为true？  
一种情况，X的左子树和右子树中均存在P或Q中的一个，如果左子树包含的是P节点，那么右子树只能包含Q节点。 即f(x左子树) && f(x右子树)  
另一种情况，X本身是P或Q，它的子树中存在另一个。 即（X = P || X = Q ） && (f(x左子树)!=nil || f(x右子树) != nil)  
刚开始无法理解为什么 f(x左子树) && f(x右子树)，实际上是审题出现了问题。  
题目中说明了，指定两个节点，证明这个节点一定存在，不存在不存在的情况。如果在左子树中存在，且在右子树中存在，那么当前的就是公共祖先，
所以return root。  

2、从前序与中序构建二叉树  
递归：  
最主要的是理解几个重要位置
假设i是中序遍历的根节点。  
那么在Go语言中inOrder[:i]就是左子树，inOrder[i+1:]就是右子树。  
在preOrder中，preOrder[1:inOrder[:i]的长度+1]就是左子树。preOrder[inOrder[:i]的长度+1]:]就是右子树。  
为什么长度需要加1？  
在左子树中，根据Go切片的定义，取部分切片时，：左右两边为前开后闭区间，即[)。  
另一个需要注意的地方是  
var root = &TreeNode{前序起点，nil,nil}
取TreeNode的地址。  
构建左右子树
root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])


分治算法：  
最主要的是找重复性，可能是最近，可能是最优。大问题都是由许多小的问题组成的。  
自顶而下。
 