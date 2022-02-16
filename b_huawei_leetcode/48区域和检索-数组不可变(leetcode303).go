package b_huawei_leetcode

type NumArray struct {
	nums []int
}

func ConstructorV1(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > right {
		return 0
	}
	var sum int
	for k := left; k <= right; k++ {
		if k > len(this.nums) {
			sum += 0
		} else {
			sum += this.nums[k]
		}
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func ConStructorV2(nums []int) NumArray {
	var preSum = make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{nums: preSum}
}

func (this *NumArray)SumRangeV2(left int, right int) int {
	if left > right {
		return 0
	}
	return this.nums[right + 1] - this.nums[left]
}
