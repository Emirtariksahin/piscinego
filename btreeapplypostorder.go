package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Recursively apply postorder traversal to the left and right subtrees
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)

	// Apply the given function to the current node's data
	f(root.Data)
}
