package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			node := curr.Left
			temp := node
			for temp.Right != nil {
				temp = temp.Right
			}
			temp.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}
