package leetcode

import (
	"fmt"
	"testing"
)

func Test_possibleBipartition(t *testing.T) {
	res := possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}})
	fmt.Println(res)
}
