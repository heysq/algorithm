package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	var sum int
	convertHelper(root, &sum)
	return root
}

func convertHelper(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	convertHelper(root.Right, sum)
	*sum = *sum + root.Val
	root.Val = *sum
	convertHelper(root.Left, sum)
}
