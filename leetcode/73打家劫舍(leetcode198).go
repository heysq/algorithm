package leetcode

// func rob(nums []int) int {
// 	var temp = make(map[int]int)
// 	res := robDP(nums, 0, temp)
// 	return res
// }

// func robDP(nums []int, start int, data map[int]int) int {
// 	if start >= len(nums) {
// 		return 0
// 	}
// 	if val, ok := data[start]; ok {
// 		return val
// 	}
// 	res := maxInt(robDP(nums, start+1, data), robDP(nums, start+2, data)+nums[start])
// 	data[start] = res
// 	return res

// }

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	dp_i_1, dp_i_2, dp_i := 0, 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp_i = maxInt(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}
