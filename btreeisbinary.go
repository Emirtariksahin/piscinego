package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTUtil(root, nil, nil)
}

func isBSTUtil(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}

	// Check if the current node's value is within the allowed range
	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}

	// Recursively check the left and right subtrees with updated ranges
	return isBSTUtil(node.Left, min, node) && isBSTUtil(node.Right, node, max)
}
