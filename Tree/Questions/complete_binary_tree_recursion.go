package questions

func (t *Tree) IsCompleteTree2() bool {
	count := t.NumNodes()
	return isCompleteTreeUtil(t.root, 0, count)
}

func isCompleteTreeUtil(curr *Node, index int, count int) bool {
	if curr == nil {
		return true
	}
	return index < count && isCompleteTreeUtil(curr.left, 2*index+1, count) && isCompleteTreeUtil(curr.right, 2*index+2, count)
}
