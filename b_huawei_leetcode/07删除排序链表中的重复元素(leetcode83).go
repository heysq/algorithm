package b_huawei_leetcode

/*
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
1. node 与 node.Next 值相同的时候，node.Next = node.Next.Next
2. 递归遍历，head.Next.Val == head.Val 返回head.Next
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	var p, q = head, head.Next
// 	for q != nil {
// 		if p.Val!= q.Val {
// 			p.Next = q
// 			p = q
// 		}
//         q = q.Next
//         p.Next = nil
// 	}

// 	return head
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	var p = head
// 	for p.Next != nil {
// 		if p.Next.Val == p.Val {
// 			p.Next = p.Next.Next
// 		} else {
// 			p = p.Next
// 		}
// 	}
// 	return head
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	head.Next = deleteDuplicates(head.Next)
// 	if head.Next.Val == head.Val {
// 		return head.Next
// 	}
// 	return head
// }
