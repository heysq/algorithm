package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func preorderTraversal(root *TreeNode) []int {
// 	var arr []int
// 	preAccessTree(root, &arr)
// 	return arr
// }

// func preAccessTree(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	*arr = append(*arr, root.Val)
// 	preAccessTree(root.Left, arr)
// 	preAccessTree(root.Right, arr)
// }
/*
1. 先序遍历，先将自己存入结果数组，然后一路遍历自己的左子树的左子树
2. 最左节点找到后，遍历右节点
*/
func preorderTraversal(root *TreeNode) []int {
	var arr []int
	stack := &BinaryTreeStack{}
	for root != nil || !stack.Empty() {
		for root != nil {
			arr = append(arr, root.Val)
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop()
		root = root.Right

	}
	return arr
}
