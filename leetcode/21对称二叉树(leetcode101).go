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
1. 递归检查左子树的左子树和右子树的右子树
2. 递归检查左子树的右子树和右子树的左子树
*/
// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return isSymmetricCheck(root.Left, root.Right)
// }

// func isSymmetricCheck(left *TreeNode, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	}

// 	if left == nil || right == nil {
// 		return false
// 	}

// 	if left.Val != right.Val {
// 		return false
// 	}

// 	return isSymmetricCheck(left.Left, right.Right) && isSymmetricCheck(left.Right, right.Left)
// }

/*
1. 左子树和右子树入栈
2. 每次从栈弹出两个成对的节点
3. 入栈时也是成对的节点（左子树的右子树和右子树的左子树，左子树的左子树和右子树的右子树）
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	if root.Right.Val != root.Left.Val {
		return false
	}
	queue := &BinaryTreeQueue{}
	queue.Push(root.Left)
	queue.Push(root.Right)
	for !queue.Empty() {
		u := queue.Pull()
		v := queue.Pull()
		if u == nil && v == nil {
			continue
		}

		if u == nil || v == nil {
			return false
		}

		if u.Val != v.Val {
			return false
		}
		queue.Push(u.Left)
		queue.Push(v.Right)
		queue.Push(u.Right)
		queue.Push(v.Left)
	}
	return true
}
