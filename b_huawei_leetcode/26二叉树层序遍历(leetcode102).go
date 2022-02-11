package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	var arr [][]int
	queue := &BinaryTreeQueue{}
	queue.Push(root)
	for !queue.Empty() {
		var temp []int
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Pull()
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		arr = append(arr, temp)
	}
	return arr
}
