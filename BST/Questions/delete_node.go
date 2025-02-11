package questions

func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = DeleteNode(node.left, value)
	} else if value > node.value {
		node.right = DeleteNode(node.right, value)
	} else {
		// Node with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		// Node with two children
		temp := FindMin(node.right)
		node.value = temp.value
		node.right = DeleteNode(node.right, temp.value)
	}
	return node
}
