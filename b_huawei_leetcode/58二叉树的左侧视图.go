package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leftSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var res = []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		level := queue
		res = append(res, queue[0].Val)
		queue = []*TreeNode{}
		for i:=0; i < len(level); i++{
			if level[i].Left != nil {
				queue = append(queue, level[i].Left)
			}
			if level[i].Right != nil {
				queue = append(queue, level[i].Right)
			}
		}
	}
	return res
}
