package leetcode

import (
	"algorithm/utils"
	"fmt"
	"time"
)

// 递归解法
// var TempMap = make(map[int]int)
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	if value, ok := TempMap[n]; ok {
// 		return value
// 	}
// 	result := climbStairs(n-1) + climbStairs(n-2)
// 	TempMap[n] = result
// 	return result
// }

// 非递归解法
// 递归公示 f(n) = f(n-1) + f(n-2)
// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var result, first, second = 0, 1, 2
	for i := 3; i <= n; i++ {
		result = first + second
		first = second
		second = result
	}
	return result
}

func ClimbStairs(n int) {
	defer utils.Cost(time.Now())
	result := climbStairs(n)
	fmt.Println(result)
}
