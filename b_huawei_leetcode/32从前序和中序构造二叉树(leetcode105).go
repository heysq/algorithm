package b_huawei_leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootValue := preorder[preStart]
	var rootIndex int = -1
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootValue {
			rootIndex = i
			break
		}
	}
	if rootIndex == -1 {
		return nil
	}
	leftSize := rootIndex - inStart
	root := &TreeNode{Val: rootValue}
	root.Left = build(preorder, inorder, preStart+1, preStart+leftSize, inStart, rootIndex-1)
	root.Right = build(preorder, inorder, preStart+leftSize+1, preEnd, rootIndex+1, inEnd)
	return root
}
