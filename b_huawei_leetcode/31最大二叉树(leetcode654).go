package b_huawei_leetcode

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	return construct(nums, 0, len(nums)-1)
	// root.Left =
}

func construct(nums []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}

	var maxVal, maxIndex = nums[i], math.MinInt
	fmt.Println(i,j)
	for k := i; k <= j; k++ {
		if nums[k] > maxVal {
			maxVal = nums[k]
			maxIndex = k
		}
	}
	root := &TreeNode{Val: maxVal}
	root.Left = construct(nums, i, maxIndex-1)
	root.Right = construct(nums, maxIndex+1, j)
	return root
}
