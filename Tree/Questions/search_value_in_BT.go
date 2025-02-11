package questions

func (t *Tree) SearchBT(value int) bool {
	return searchBT(t.root, value)
}

func searchBT(root *Node, value int) bool {
	//Implement your solution here
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	}
	return false || searchBT(root.left, value) || searchBT(root.right, value)
}
