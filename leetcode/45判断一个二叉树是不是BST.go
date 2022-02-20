package leetcode

func IsBST(root *TreeNode) bool {
	return IsBSTHelper(root, nil, nil)
}

// min < root < max
func IsBSTHelper(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}

	if max != nil && root.Val >= min.Val {
		return false
	}

	return IsBSTHelper(root.Left, min, root) && IsBSTHelper(root, root, max)
}
