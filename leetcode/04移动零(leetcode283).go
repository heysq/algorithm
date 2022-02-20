package leetcode

import (
	"algorithm/utils"
	"fmt"
	"time"
)

/*
https://leetcode-cn.com/problems/move-zeroes/
变量 i 永远记录的是0的位置
用变量j循环的遍历数组
遍历完成后，将i到最后的元素都设置为0
*/
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
