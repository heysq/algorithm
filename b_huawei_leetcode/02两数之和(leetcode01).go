package b_huawei_leetcode

import (
	"algorithm/utils"
	"fmt"
	"time"
)

/*
map 存储遍历过的索引
遇到存在的字段，return
https://leetcode-cn.com/problems/two-sum
*/
func twoSum(nums []int, target int) []int {
	var result = []int{}
	var TempMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, ok := TempMap[target-nums[i]]; ok {
			return []int{value, i}
		}
		TempMap[nums[i]] = i
	}
	return result
}

func TwoSum(nums []int, target int) {
	defer utils.Cost(time.Now())
	result := twoSum(nums, target)
	fmt.Println(result)
}
