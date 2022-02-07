package b_huawei_leetcode

import (
	"algorithm/utils"
	"fmt"
	"time"
)

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var i, j int
	for ; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}

func MoveZeroes(nums []int) {
	defer utils.Cost(time.Now())
	moveZeroes(nums)
	fmt.Println(nums)
}
