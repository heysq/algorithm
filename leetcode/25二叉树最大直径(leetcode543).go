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
利用二叉树的后续遍历思想
获得左右子树的最大直径
然后记录遍历途中的最大直径
*/
func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int
	depth(root, &maxDiameter)
	return maxDiameter
}

func depth(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left, maxDiameter)
	rightDepth := depth(root.Right, maxDiameter)
	if leftDepth+rightDepth > *maxDiameter {
		*maxDiameter = leftDepth + rightDepth
	}
	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}
