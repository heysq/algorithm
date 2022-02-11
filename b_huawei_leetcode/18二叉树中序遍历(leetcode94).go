package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func inorderTraversal(root *TreeNode) []int {
// 	var arr []int
// 	accessTree(root, &arr)
// 	return arr
// }

// func accessTree(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	accessTree(root.Left, arr)
// 	*arr = append(*arr, root.Val)
// 	accessTree(root.Right, arr)
// }

func inorderTraversal(root *TreeNode) []int {
	var arr []int
	stack := &BinaryTreeStack{}
	for root != nil || !stack.Empty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop()
		arr = append(arr, root.Val)
		root = root.Right
	}
	return arr
}
