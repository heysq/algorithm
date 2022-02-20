package leetcode

/*
https://leetcode-cn.com/problems/linked-list-cycle-ii/
1. 判断链表有环
2. 快指针和慢指针相遇的时候，其中一个回到head，每次走一步，再次相遇就是环的交点
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var fastPtr, slowPtr = head, head
	var loopExists bool
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
		if fastPtr == slowPtr {
			loopExists = true
			break
		}
	}
	if loopExists {
		slowPtr = head
		for slowPtr != fastPtr {
			slowPtr = slowPtr.Next
			fastPtr = fastPtr.Next
		}
		return slowPtr
	}
	return nil
}
