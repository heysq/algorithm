package b_huawei_leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	preValue := nums[0]
	var result = preValue
	for i := 1; i < len(nums); i++ {
		preValue = preValue + nums[i]
		if nums[i] > preValue {
			preValue = nums[i]
		}
		if result < preValue {
			result = preValue
		}
	}
	return result
}
