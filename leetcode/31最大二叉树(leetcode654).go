package leetcode

import (
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

/*
每次循环取出最大值，做为根节点
1、最大值索引左边的是左子树
2. 最大值索引右边的是右子树
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
