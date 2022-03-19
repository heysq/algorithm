package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
利用后续遍历，知道左右子树的最大深度
*/
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	left := maxDepth(root.Left)
// 	right := maxDepth(root.Right)
// 	depth := math.Max(float64(left), float64(right)) + 1
// 	return int(depth)
// }

func maxDepth(root *TreeNode) int {
	var max, depth int
	traverse(root, &depth, &max)
	return max
}

func traverse(node *TreeNode, depth *int, max *int) {
	if node == nil {
		if *depth > *max {
			*max = *depth
		}
		return
	}
	*depth++
	traverse(node.Left, depth, max)
	traverse(node.Right, depth, max)
	*depth--
}
