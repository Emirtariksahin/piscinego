package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil || root == nil {
		return root
	}

	// Update the parent pointer of the replaced subtree
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	// Update the parent pointer of the replacing subtree
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
