package leetcode

import (
	"sort"
)

// 两数之和升级
// 返回所有可能的和，如果数组重复，可以跳过相同的，优化时间复杂度
func twoSumHighLevel(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var low, high = 0, len(nums) - 1
	for low < high {
		sum := nums[low] + nums[high]
		left, right := nums[low], nums[high]
		if sum == target {
			res = append(res, []int{nums[low], nums[high]})
			for low < high && left == nums[low] {
				low++
			}
			for low < high && right == nums[high] {
				high--
			}
		} else if sum < target {
			for low < high && left == nums[low] {
				low++
			}
		} else if sum > target {
			for low < high && right == nums[high] {
				high++
			}
		}
	}
	return res
}
