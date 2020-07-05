学习笔记  
#第二周学习笔记  

##本周学习要点  
哈希表：又散列表。通过散列函数将要存储的值映射到一个位置，如果发生哈希碰撞，可以拉出个链表。常用于redis。  
平均时间复杂度：插入，删除，查找均为O(1)  
红黑树的平均时间复杂度都为O(logN)  
  
二叉搜索树：又叫有序二叉树  
左子树上所有节点的值均小于根节点的值  
右子树上所有节点的值均大于根节点的值  
其中序遍历为升序排列。  
查询和插入的平均时间复杂度为O(logN)  
  
堆：常用于求topK，利用大顶堆或小顶堆。  
平均时间复杂度：查找O(1),插入O(logN),删除O(logN)。最好的实现方式是严格的斐波那契堆。  

二叉堆是完全二叉树来实现的。  
1、是一棵完全二叉树  
2、树中任意节点的值总是>=其子节点的值。  

通过2*i+1找到左子节点 2*i+2找到右子节点。  
建立堆必须自上而下的顺序执行堆化。  



Operation	    find-min	delete-min	insert	  decrease-key	meld  
Binary	        Θ(1)	    Θ(log n)	O(log n)	O(log n)	Θ(n)  
Leftist	        Θ(1)	    Θ(log n)	Θ(log n)	O(log n)	Θ(log n)  
Binomial	    Θ(1)	    Θ(log n)	Θ(1)	    Θ(log n)	O(log n)  
Fibonacci	    Θ(1)	    O(log n)	Θ(1)	    Θ(1)	    Θ(1)  
Pairing	        Θ(1)	    O(log n)	Θ(1)	    O(log n)	Θ(1)  
Brodal	        O(1)	    O(log n)	Θ(1)	    Θ(1)	    Θ(1)  
Rank-pairing	Θ(1)	    O(log n)	Θ(1)	    Θ(1)	    Θ(1)  
Strict FibonacciΘ(1)	    O(log n)	Θ(1)	    Θ(1)	    Θ(1)  
2-3 heap	    O(log n)	O(log n)	O(log n)	Θ(1)	     ?  



##作业
语言：Go
1、HashMap  
Go语言中，HashMap的实现是基于map的实现。  
在hashmap的package中，在Map结构体中，封装了map变量。添加和获取即在map中新增和获取，是非线程安全的。  

Go语言实现heap，位于package container/heap中。实现了以下方法：    
func Fix(h Interface, i int)  
func Init(h Interface)  
func Pop(h Interface) interface{}  
func Push(h Interface, x interface{})  
func Remove(h Interface, i int) interface{}  

func Fix(h Interface, i int)用于当index为i的位置值变更时，重建堆。调用此方法时间复杂度比remove低，为O(logN)。  
