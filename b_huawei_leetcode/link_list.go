package b_huawei_leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkList(elems []int) *ListNode {
	var head *ListNode = &ListNode{}
	var arr []*ListNode
	for _, elem := range elems {
		arr = append(arr, &ListNode{Val: elem})
	}
	head.Next = arr[0]
	for i := 0; i < len(elems)-1; i++ {
		arr[i].Next = arr[i+1]
	}
	return head.Next
}

func PrintLinkList(head *ListNode) {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	fmt.Println(arr)
}
