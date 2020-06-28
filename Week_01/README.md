学习笔记
#第一周学习笔记
***
##1、心得体会
五毒神掌，贵在坚持。  
由于工作加班，日常时间不够，有的题解未完全做到所有解法尝试一遍，希望下周改进。    
算法优化的本质，空间换时间。
一维解决不了的问题，上升到二维试试。


##2、本周学习要点
###2.1 数组、链表、跳表
时间复杂度：  
数组 插入 O(n)   删除 O(n)   查找 O(1)  
链表 插入 O(1)   删除 O(1)   查找 O(n)   
跳表 插入 O(logN)   删除 O(logN)  查找 O(logN)

跳表：链表有序的情况，对标平衡二叉树和二分查找。原理简单，方便扩展，容易实现，效率更高。redis中有使用。


###2.2 栈、队列、优先队列、双端队列
栈Stack：FILO （先入后出） 添加、删除 O(1) 查询 O(n)  
队列Queue：FIFO （先进先出） 添加、删除 O(1)  查询 O(n)  
双端队列Double-End Queue： 队首可pop和push 队尾可pop和push 添加、删除 O(1)  查询 O(n)  
优先队列PriorityQueue： 添加O(1) 取出O(logN)按元素优先级取出  底层实现数据结构复杂多样，可用heap、二叉搜索树（BST）、treap实现。


##3、分析 Queue 和 Priority Queue 的源码
语言：Go  
链接：go-datastructures: github.com/golang-collections/go-datastructures/queue

Go语言中，Queue和PriorityQueue均在queue的package中。  
都实现了：dispose()、disposed()、empty()、get(number int)、len()、put(items...)方法  

Queue文件：go-datastructures/queue/queue.go   
链接：https://github.com/golang-collections/go-datastructures/blob/master/queue/queue.go  
PriorityQueue文件：go-datastructures/queue/priority_queue.go   
链接：https://github.com/golang-collections/go-datastructures/blob/master/queue/priority_queue.go


func (q *Queue) Dispose()  
dispose队列，后续该队列调用get或put都将返回error。    
queue.disposed设置为true，队列清空  
func (pq *PriorityQueue) Dispose()  
避免对该队列的任何读写操作，释放资源。
pq.disposed设置为true，队列清空   

func (q *Queue) Disposed() bool  
确认是否已被dispose,返回q.disposed  
func (pq *PriorityQueue) Disposed() bool  
确认是否已被dispose,返回pq.disposed

func (q *Queue) Empty() bool  
判空 返回item长度是否为0  
func (pq *PriorityQueue) Empty() bool  
判空 返回item长度是否为0  


func (q *Queue) Get(number int64) ([]interface{}, error)      
如果number<1 返回空切片；如果已经被dispose，返回disposedError；如果队列items是空，创建一个新的newSema，放到waiter中等数据，
调用items的get方法，返回number数量的数据。  
func (pq *PriorityQueue) Get(number int) ([]Item, error)
同上

func（q * Queue）Len（）int64  
返回队列items长度
func（pq * PriorityQueue）Len（）int
同上

func (q *Queue) Put(items ...interface{}) error  
向队列中添加指定元素  向q.items中append元素。 这块没搞懂。 
`for {   
  		sema := q.waiters.get()    
  		if sema == nil {  
  			break
  		}
  		sema.response.Add(1)
  		sema.wg.Done()
  		sema.response.Wait()
  		if len(q.items) == 0 {
  			break
  		}
  	}`
  	

同时：
`func (w *waiters) get() *sema {
 	if len(*w) == 0 {
 		return nil
 	}
 	sema := (*w)[0]
 	copy((*w)[0:], (*w)[1:])
 	(*w)[len(*w)-1] = nil // or the zero value of T
 	*w = (*w)[:len(*w)-1]
 	return sema
 }`

func (pq *PriorityQueue) Put(items ...Item) error
同上  


func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error)  
返回与checker方法相匹配的items

func (pq *PriorityQueue) Peek() Item  
不移除元素的条件下，查询下一个元素。返回items[0]