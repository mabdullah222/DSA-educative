package questions

func (t *Tree) TreeToListRec() *Node {
	return treeToListRec(t.root)
}
func treeToListRec(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	var Head, Tail, tempHead *Node

	if curr.left == nil && curr.right == nil {
		curr.left = curr
		curr.right = curr
		return curr
	}

	if curr.left != nil {
		Head = treeToListRec(curr.left)
		Tail = Head.left
		curr.left = Tail
		Tail.right = curr
	} else {
		Head = curr
	}

	if curr.right != nil {
		tempHead = treeToListRec(curr.right)
		Tail = tempHead.left
		curr.right = tempHead
		tempHead.left = curr
	} else {
		Tail = curr
	}
	Head.left = Tail

	Tail.right = Head
	return Head
}
