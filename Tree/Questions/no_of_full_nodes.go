package questions

func (t *Tree) NumFullNodesBT() int {
	return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *Node) int {
	if curr == nil {
		return 0
	}
	var count int
	if curr.left != nil && curr.right != nil {
		count = 1
	}
	count += numFullNodesBT(curr.left)
	count += numFullNodesBT(curr.right)
	return count
}
