package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

/*func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	} else {
		if data < root.Data {
			root.Left = BTreeInsertData(root.Left, data)
			root.Left.Parent = root
		} else {
			root.Right = BTreeInsertData(root.Right, data)
			root.Right.Parent = root
		}
		return root
	}
}
*/
