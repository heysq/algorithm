package leetcode

import "sort"

func nSum(nums []int, target int, n int) [][]int {
	sort.Ints(nums)
	res := nSumHelper(nums, target, n, 0)
	return res
}

func nSumHelper(nums []int, target int, n int, start int) [][]int {
	size := len(nums)
	res := [][]int{}
	if n < 2 || size < n {
		return res
	}
	if n == 2 {
		low, high := start, size-1
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
	} else if n > 2 {
		for i := start; i < n; i++ {
			_res := nSumHelper(nums, target-nums[i], n-1, i+1)
			for j := 0; j < len(_res); j++ {
				_res[j] = append(_res[j], nums[i])
				res = append(res, _res[j])
			}
			for i < size-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
