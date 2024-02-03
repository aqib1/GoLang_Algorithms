package bst

// Time complexity O(n) and space complexity O(n)
func isSymmetric(root *TreeNode) bool {
	return isSymmetricRec(root, root)
}

func isSymmetricRec(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val &&
		isSymmetricRec(left.Left, right.Right) &&
		isSymmetricRec(left.Right, right.Left)
}
