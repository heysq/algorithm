package leetcode

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	if len(postorder) == 1 {
		return true
	}
	root := postorder[len(postorder)-1]
	var left int
	for left < len(postorder) {
		if postorder[left] < root {
			left++
		} else {
			break
		}
	}
	// [5, 2, -17, -11, 25, 76, 62, 98, 92, 61]
	for j := left; j < len(postorder); j++ {
		if postorder[j] < root {
			fmt.Println(j)
			return false
		}
	}
	return verifyPostorder(postorder[:left]) && verifyPostorder(postorder[left:len(postorder)-1])
}
