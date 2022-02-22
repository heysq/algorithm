package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func robIII(root *TreeNode) int {
	hash := make(map[*TreeNode]int)
	return robHelper(root, hash)
}
func robHelper(root *TreeNode, hash map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if val, ok := hash[root]; ok {
		return val
	}
	// 打劫 root
	var doVal = root.Val
	if root.Left != nil {
		doVal = doVal + robHelper(root.Left.Left, hash) + robHelper(root.Left.Right, hash)
	}
	if root.Right != nil {
		doVal = doVal + robHelper(root.Right.Left, hash) + robHelper(root.Right.Right, hash)
	}

	var notDoVal = robHelper(root.Left, hash) + robHelper(root.Right, hash)
	res := maxInt(doVal, notDoVal)
	hash[root] = res
	return res

}
