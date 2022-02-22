package leetcode

func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return maxInt(robIIHelper(nums, 0, len(nums)-2), robIIHelper(nums, 1, len(nums)-1))
}

func robIIHelper(nums []int, start, end int) int {
	var dp_i, dp_i_1, dp_i_2 int
	for i := end; i >= start; i-- {
		dp_i = maxInt(dp_i_1, dp_i_2+nums[i])
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}
