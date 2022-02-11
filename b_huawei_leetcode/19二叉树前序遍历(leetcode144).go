package b_huawei_leetcode

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
