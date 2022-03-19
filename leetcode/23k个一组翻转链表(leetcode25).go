package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseAB(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseAB(a *ListNode, b *ListNode) *ListNode {
	if a == nil || a.Next == nil {
		return a
	}
	var pre, cur, next *ListNode = a, a, nil
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
