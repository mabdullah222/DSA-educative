package questions

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func isEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if (node1 == nil && node2 != nil) || (node2 == nil && node1 != nil) {
		return false
	}
	if node1.value != node2.value {
		return false
	}
	return true && isEqual(node1.left, node2.left) && isEqual(node1.right, node2.right)
}
