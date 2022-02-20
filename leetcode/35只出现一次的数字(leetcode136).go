package leetcode

func singleNumber(nums []int) int {
	var result int

	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}
