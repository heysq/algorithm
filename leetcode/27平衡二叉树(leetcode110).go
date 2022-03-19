package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
 二叉树后续遍历思想
 1. 左树是不是平衡二叉树
 2. 右树是不是平衡二叉树
 4. 自己是不是平衡二叉树
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return balance(root) != -1
}

func balance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := balance(root.Left)
	right := balance(root.Right)
	if left == -1 || right == -1 || math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}
	return int(math.Max(float64(left), float64(right))) + 1
}
