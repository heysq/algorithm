package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil && right == nil {
		return nil
	}

	if left == nil {
		return right
	} else {
		return left
	}
}

func lowestV2(root, p, q *TreeNode) *TreeNode {
	var hash = make(map[*TreeNode]*TreeNode, 0)
	lowestHelper(root, hash)
	if hash[p] == nil || hash[q] == nil {
		return root
	}
	var tempMap = make(map[*TreeNode]struct{})
	for node := hash[p]; node != nil; node = hash[node] {
		tempMap[node] = struct{}{}
	}
	for node := hash[q]; node != nil; node = hash[node] {
		if _, ok := tempMap[node]; ok {
			return node
		}
	}
	return nil
}

func lowestHelper(root *TreeNode, hash map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}
	if _, ok := hash[root]; !ok {
		hash[root] = nil
	}
	if root.Left != nil {
		hash[root.Left] = root
	}
	if root.Right != nil {
		hash[root.Right] = root
	}
	lowestHelper(root.Left, hash)
	lowestHelper(root.Right, hash)
}
