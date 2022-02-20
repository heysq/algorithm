package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
1. 两个指针不相等的情况下，每次其中一个为nil的时候，切换为另一个的开头，相当于补全两个链表到同一个长度遍历
2. 遍历一遍2个链表获取长度，让短的那个先走diff个长度，然后同时遍历，有相等的点就是相交
*/
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	if headA == nil || headB == nil {
// 		return nil
// 	}
// 	var aPtr, bPtr = headA, headB
// 	for aPtr != bPtr {
// 		if aPtr == nil {
// 			aPtr = headB
// 		} else {
// 			aPtr = aPtr.Next
// 		}

// 		if bPtr == nil {
// 			bPtr = headA
// 		} else {
// 			bPtr = bPtr.Next
// 		}
// 	}
// 	return aPtr
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var aLen, bLen int
	var aPtr, bPtr = headA, headB
	for aPtr != nil {
		aLen++
		aPtr = aPtr.Next
	}

	for bPtr != nil {
		bLen++
		bPtr = bPtr.Next
	}

	var head = headA
	diff := aLen - bLen
	if aLen < bLen {
		head = headB
		diff = bLen - aLen
	}

	for i := 0; i < diff; i++ {
		head = head.Next
	}

	if aLen < bLen {
		headB = head
	} else {
		headA = head
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
