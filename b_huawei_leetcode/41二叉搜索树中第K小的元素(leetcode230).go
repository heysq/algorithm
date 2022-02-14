package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var rank, result int
	kthSmall(root, k, &rank, &result)
	return result
}

func kthSmall(root *TreeNode, k int, rank, result *int ) {
	if root == nil {
		return
	}
	kthSmall(root.Left, k, rank, result)
	*rank++
	if *rank == k {
		*result = root.Val
		return
	}
	kthSmall(root.Right, k, rank, result)
}
