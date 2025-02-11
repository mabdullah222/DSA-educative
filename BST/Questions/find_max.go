package questions

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root
	for node.right != nil {
		node = node.right
	}
	//Implement your solution here

	return node
}
