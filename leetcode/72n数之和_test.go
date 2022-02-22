package leetcode

import (
	"fmt"
	"testing"
)

func Test_nSum(t *testing.T) {
	res := nSum([]int{-1, 0, 1, 2, -1, -4}, 0, 3)
	fmt.Println(res)
}
