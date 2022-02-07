package b_huawei_leetcode

import (
	"testing"
)

func TestNewLinkList(t *testing.T) {
	head := NewLinkList([]int{1, 2, 3, 4, 5, 6})
	PrintLinkList(head)
}
