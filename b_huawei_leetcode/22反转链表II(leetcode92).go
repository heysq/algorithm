package b_huawei_leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	if left == 1 {
// 		return reverseN(head, right)
// 	}
// 	head.Next = reverseBetween(head.Next, left-1, right-1)
// 	return head
// }

// var successor *ListNode = nil

// func reverseN(head *ListNode, n int) *ListNode {
// 	if n == 1 {
// 		successor = head.Next
// 		return head
// 	}
// 	last := reverseN(head.Next, n-1)
// 	head.Next.Next = head
// 	head.Next = successor
// 	return last
// }

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = head
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	var curr *ListNode = pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next	
	}
	return head
}
