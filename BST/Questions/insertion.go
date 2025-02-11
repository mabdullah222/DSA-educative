package questions

func (t *Tree) Add(value int) {
	t.root = addNode(t.root, value)
}

func addNode(n *Node, value int) *Node {
	//Implement your solution here
	if n == nil {
		newNode := new(Node)
		newNode.value = value
		return newNode
	}

	if value > n.value {
		n.right = addNode(n.right, value)
	} else {
		n.left = addNode(n.left, value)
	}
	return n

}
