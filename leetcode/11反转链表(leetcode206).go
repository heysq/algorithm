package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 1. 保存当前节点和当前节点的下一个节点
 2. 递归，相当于每个节点只需要关心自己与自己的下一个节点，自己的下一个节点的下一个节点必须，指向自己，然后自己的下一个节点的值为nil
*/
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	var curNode = head
// 	var preNode *ListNode = nil
// 	for curNode != nil {
// 		node := curNode.Next
// 		curNode.Next = preNode
// 		preNode = curNode
// 		curNode = node
// 	}
// 	return preNode
// }

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
