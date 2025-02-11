package questions

func (t *Tree) IsCompleteTree() bool {
	return isCompleteTree(t.root)
}

func isCompleteTree(root *Node) bool {
	que := new(Queue)
	var temp *Node
	var noChild = false
	if root != nil {
		que.Enqueue(root)
	}
	for que.Length() != 0 {
		temp = que.Dequeue().(*Node)
		if temp.left != nil {
			if noChild == true {
				return false
			}
			que.Enqueue(temp.left)
		} else {
			noChild = true
		}
		if temp.right != nil {
			if noChild == true {
				return false
			}
			que.Enqueue(temp.right)
		} else {
			noChild = true
		}
	}
	return true
}
