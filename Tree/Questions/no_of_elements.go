package questions

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return 1 + numNodes(curr.left) + numNodes(curr.right)
}
