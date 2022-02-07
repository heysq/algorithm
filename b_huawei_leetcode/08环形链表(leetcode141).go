package b_huawei_leetcode

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