package questions

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func copyMirrorTree(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	root := &Node{curr.value, nil, nil}
	root.left = copyMirrorTree(curr.right)
	root.right = copyMirrorTree(curr.left)
	return root
}
