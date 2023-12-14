package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref, Next: nil}

	// Case: empty list or newNode should be the new head
	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	// Traverse the list to find the insertion point
	current := l
	for current.Next != nil && data_ref > current.Next.Data {
		current = current.Next
	}

	// Insert the new node
	newNode.Next = current.Next
	current.Next = newNode

	return l
}
