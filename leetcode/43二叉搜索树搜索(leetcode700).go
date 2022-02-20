package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}

func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		deleteNode(root.Left, val)
	}
	if root.Val < val {
		deleteNode(root.Right, val)
	}

	// 没有左右节点，直接删除
	if root.Left == nil && root.Right == nil {
		return nil
	}
	// 只有一个节点 ，返回另一个节点
	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	// 获取右子树中最小的节点
	minNode := getMin(root.Right)
	root.Right = deleteNode(root.Right, minNode.Val)
	minNode.Left = root.Left
	minNode.Right = root.Right
	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	root = root.Left
	return getMin(root)
}
