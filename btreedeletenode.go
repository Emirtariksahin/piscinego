package piscine

func minValueNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: Node to be deleted has no children or one child
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	// Case 3: Node to be deleted has two children
	successor := minValueNode(node.Right)
	node.Data = successor.Data

	// Recursively delete the in-order successor
	root = BTreeDeleteNode(root, successor)
	return root
}
