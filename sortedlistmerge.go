package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Create a dummy node to simplify the code
	dummy := &NodeI{}
	current := dummy

	// Traverse both lists and merge in ascending order
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// If either list is not empty, append the remaining nodes
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	return dummy.Next
}
