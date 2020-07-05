package Week_02

/**
* Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	//1、哈希表
	var cache = make(map[*ListNode]int, 0)
	for head != nil {
		if _, ok := cache[head]; ok {
			return true
		}
		cache[head] = head.Val
		head = head.Next
	}
	return false

	//2、快慢指针
	if head == nil {
		return false
	}

	fast := head.Next
	for fast != nil && fast.Next != nil && head != nil {
		if fast.Next == head {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}
	return false
}
