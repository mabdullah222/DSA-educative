package questions

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
	if curr == nil {
		return 0
	}
	return curr.value + sumAllBT(curr.left) + sumAllBT(curr.right)
}
