package questions

func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}

func treeDepth(root *Node) int {
	if root == nil {
		return -1
	}
	left_depth := 1 + treeDepth(root.left)
	right_depth := 1 + treeDepth(root.right)
	if left_depth > right_depth {
		return left_depth
	}
	return right_depth
}
