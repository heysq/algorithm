package leetcode

import (
	"fmt"
	"testing"
)

func TestMyStack_GetNum(t *testing.T) {
	s := MyStack{}
	s.Push("3")
	s.Push("5")
	fmt.Println(s.Empty())
	fmt.Println(s.GetNum())
}

func Test_decodeString(t *testing.T) {
	result := decodeString("3[aef4[ef]r]c")
	fmt.Println(result)
}

// zzzyyyypqjkjkjkjkjkjkjkjkefef
// zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef
