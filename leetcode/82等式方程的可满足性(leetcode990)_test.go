package leetcode

import (
	"fmt"
	"testing"
)

func Test_equationsPossible(t *testing.T) {
	result := equationsPossible([]string{"b==a", "a==b"})
	fmt.Println(result)
}
