package bst

func preorderTraversal(root *TreeNode) []int {
	data := make([]int, 0, 10)
	preorder(root, &data)
	return data
}

// Time complexity O(n)
// Space complexity OLog(n)
func preorder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	*data = append(*data, root.Val)
	preorder(root.Left, data)
	preorder(root.Right, data)
}
