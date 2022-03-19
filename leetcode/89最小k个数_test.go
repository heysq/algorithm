package leetcode

import (
	"fmt"
	"testing"
)

func Test_smallestK(t *testing.T) {
	result := smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4)
	fmt.Println(result)
}
