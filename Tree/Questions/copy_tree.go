package questions

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyTree(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	newRoot := &Node{curr.value, nil, nil}
	newRoot.left = copyTree(curr.left)
	newRoot.right = copyTree(curr.right)
	return newRoot
}
