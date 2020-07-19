学习笔记  

1、思考  

贪心算法是一种思想，并非真正的算法。  
熟练掌握了dfs的模版使用，对于BFS模版的使用还需要进一步实践。  

2、算法

BFS Go语言模版：    
    func bfs(root *TreeNode, level int) {	  
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

queue 方式：

    //初始化queue
    var queue = make([]*TreeNode, 0)
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
		queue = queue[length:]`