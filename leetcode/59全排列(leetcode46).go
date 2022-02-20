package leetcode

func permute(nums []int) [][]int {
	var res [][]int
	premuteHelper(nums, &[]int{}, &res)
	return res
}

func premuteHelper(nums []int, track *[]int, res *[][]int) {
	if len(*track) == len(nums) {

		a := make([]int, len(*track))
		copy(a, *track)
		*res = append(*res, a)
		return
	}

	for i := 0; i < len(nums); i++ {
		if intArrayContain(*track, nums[i]) {
			continue
		}
		*track = append(*track, nums[i])
		premuteHelper(nums, track, res)
		*track = (*track)[:len(*track)-1]
	}
}

func intArrayContain(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
