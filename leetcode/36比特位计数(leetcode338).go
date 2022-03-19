package leetcode

import "fmt"

// func countBits(n int) []int {
// 	var nums = make([]int, n+1)
// 	for i := 0; i <= n; i++ {
// 		for x := i; x > 0; x = x & (x - 1) {
// 			nums[i]++ // 记录操作次数
// 		}
// 	}
// 	return nums
// }

func countBits(n int) []int {
	var nums = make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = nums[i&(i-1)] + 1 // nums[2] = nums[2 &(2-1)] + 1
		fmt.Println(nums[i])
	}
	return nums
}
