package b_huawei_leetcode

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