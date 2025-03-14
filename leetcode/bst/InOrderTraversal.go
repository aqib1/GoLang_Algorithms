package bst

func inorderTraversal(root *TreeNode) []int {
	data := make([]int, 0, 10)
	inorder(root, &data)
	return data
}

// Time complexity O(n)
// Space complexity O(n)
func inorder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, data)
	*data = append(*data, root.Val)
	inorder(root.Right, data)
}
