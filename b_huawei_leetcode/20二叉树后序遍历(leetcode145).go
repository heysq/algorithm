package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func postorderTraversal(root *TreeNode) []int {
// 	var arr []int
// 	postorderAccessTree(root, &arr)
// 	return arr
// }

// func postorderAccessTree(root *TreeNode, arr *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	postorderAccessTree(root.Left, arr)
// 	postorderAccessTree(root.Right, arr)
// 	*arr = append(*arr, root.Val)
// }

func postorderTraversal(root *TreeNode) []int {
	var arr []int
	stack := &BinaryTreeStack{}
	var preNode *TreeNode = nil
	for root != nil || !stack.Empty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop()
		if root.Right == nil || root.Right == preNode {
			arr = append(arr, root.Val)
			preNode = root
			root = nil
		} else {
			stack.Push(root)
			root = root.Right
		}
	}
	return arr
}
