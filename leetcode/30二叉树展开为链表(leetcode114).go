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
找到节点左子树的最右侧节点temp
 然后节点右子树挂到这个节点的右子树的下边
 curr节点的右子树设置为curr的左子树
 curr的左子树设置为nil

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
