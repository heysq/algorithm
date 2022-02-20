package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree1(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build1(inorder, postorder []int, inStart, inEnd, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	rootVal, rootIndex := postorder[postEnd], -1
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	leftSize := rootIndex - inStart
	root := &TreeNode{Val: rootVal}
	root.Left = build1(inorder, postorder, inStart, rootIndex-1, postStart, postStart+leftSize-1)
	root.Right = build1(inorder, postorder, rootIndex+1, inEnd, postStart+leftSize, postEnd-1)
	return root
}
