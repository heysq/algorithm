package b_huawei_leetcode

/*
https://leetcode-cn.com/problems/linked-list-cycle/
1. map存储遍历过的节点，循环在map中取到标记，则链表存在环
2. 快慢指针，如果快指针与慢指针值相同，则存在环，fast.Next != nil && fast.Next.Next != nil
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func hasCycle(head *ListNode) bool {
// 	if head == nil || head.Next == nil {
// 		return false
// 	}
// 	var tempMap = make(map[*ListNode]struct{})    
// 	for head != nil {
// 		if _, ok := tempMap[head]; ok {
// 			return true
// 		}
// 		tempMap[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return false
// }

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var fastPtr, slowPtr = head, head
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
		if fastPtr == slowPtr {
			return true
		}
	}
	return false
}