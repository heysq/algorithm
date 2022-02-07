package b_huawei_leetcode

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
