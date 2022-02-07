package b_huawei_leetcode

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	head := NewLinkList([]int{1,1,2,1})
	fmt.Println(isPalindrome(head))
}
