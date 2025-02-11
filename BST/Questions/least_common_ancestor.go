package questions

func (t *Tree) LcaBST(first int, second int) (int, bool) {
	if t.root == nil {
		return 0, false
	}
	return LcaBST(t.root, first, second)
}

func LcaBST(curr *Node, first int, second int) (int, bool) {
	if curr.value > first && curr.value > second {
		return LcaBST(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return LcaBST(curr.right, first, second)
	}
	return curr.value, true
}
