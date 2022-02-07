package b_huawei_leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var curNode = head
	var preNode *ListNode = nil
	for curNode != nil {
		node := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = node
	}
	return preNode
}
