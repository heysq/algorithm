package leetcode

/*
将链表折半，找到链表的中点，翻转中点后的链表，然后循环判断
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	halfNode := getHalfNode(head)
	reverse := reverseList(halfNode.Next)
	for reverse != nil {
		if reverse.Val != head.Val {
			return false
		}
		reverse = reverse.Next
		head = head.Next
	}
	return true
}

func getHalfNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
