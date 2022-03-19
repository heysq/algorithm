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
1. 前序和后续遍历无法唯一确定二叉树的结构
2. 前序遍历可以确定根节点，然后更几点的下一个节点左子树
3. 递归拆分数组 
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return build2(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1)
}

func build2(preorder, postorder []int, preStart, preEnd, postStart, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{Val: preorder[preStart]}
	}
	rootVal := preorder[preStart]
	rootLeftVal := preorder[preStart+1]
	leftIndex := -1
	for i := postStart; i <= postEnd; i++ {
		if postorder[i] == rootLeftVal {
			leftIndex = i
			break
		}
	}
	leftSize := leftIndex - postStart + 1
	root := &TreeNode{Val: rootVal}
	root.Left = build2(preorder, postorder, preStart+1, preStart+leftSize, postStart, leftIndex)
	root.Right = build2(preorder, postorder, preStart+leftSize+1, preEnd, leftIndex+1, postEnd-1)
	return root
}
