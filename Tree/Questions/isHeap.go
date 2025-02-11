package questions

func (t *Tree) IsHeap() bool {
	parentValue := -99999999
	return t.IsCompleteTree() && isHeapUtil(t.root, parentValue)
}

func isHeapUtil(curr *Node, parentValue int) bool {
	//Implement your solution here
	if curr == nil {
		return true
	}
	return parentValue < curr.value && isHeapUtil(curr.left, curr.value) && isHeapUtil(curr.right, curr.value)
}
