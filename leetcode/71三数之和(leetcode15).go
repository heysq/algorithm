package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		_res := twoSumNormal(nums, i+1, 0-nums[i])
		for j := 0; j < len(_res); j++ {
			_res[j] = append(_res[j], nums[i])
			res = append(res, _res[j])
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSumNormal(nums []int, start int, target int) [][]int {
	low, high := start, len(nums)-1
	res := [][]int{}
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
				high--
			}
		}
	}
	return res
}
