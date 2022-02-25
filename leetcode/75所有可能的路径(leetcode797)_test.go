package leetcode

import (
	"fmt"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	arr := [][]int{{1, 2}, {3}, {3}, {}}
	res := allPathsSourceTarget(arr)
	fmt.Println(res)
}
