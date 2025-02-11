package questions

func (t *Tree) FindMin() *Node {

	var node *Node = t.root

	//Implement your solution here
	for node.left != nil {
		node = node.left
	}
	return node
}
