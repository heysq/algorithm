package b_huawei_leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
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
