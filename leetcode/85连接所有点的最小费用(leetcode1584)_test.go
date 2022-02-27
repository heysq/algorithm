package leetcode

import (
	"fmt"
	"testing"
)

func Test_minCostConnectPoints(t *testing.T) {
	result := minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})
	fmt.Println(result)
}
