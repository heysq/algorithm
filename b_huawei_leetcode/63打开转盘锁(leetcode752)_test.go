package b_huawei_leetcode

import (
	"fmt"
	"testing"
)

func Test_openLock(t *testing.T) {
	result := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	fmt.Println(result)
}
