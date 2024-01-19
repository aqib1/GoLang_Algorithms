package bst

// Time complexity OLog(N) or O(h)
// Space complexity OLog(N) or O(h)
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return find(root, p, nil)
}

func find(root *TreeNode, p *TreeNode, successor *TreeNode) *TreeNode {
	if root == nil {
		return successor
	}

	if p.Val >= root.Val {
		return find(root.Right, p, successor)
	} else {
		return find(root.Left, p, root)
	}
}

// Time complexity OLog(N) or O(h)
// Space complexity O(1)
func findIterative(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode = nil
	for root != nil {
		if p.Val >= root.Val {
			root = root.Right
		} else {
			successor = root
			root = root.Left
		}
	}
	return successor
}
