package questions

import "fmt"

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
	//Implement your solution here
	if node == nil {
		return
	}
	nthInOrder(node.left, index, counter)
	(*counter) += 1
	if index == (*counter) {
		fmt.Print(node.value)
	}
	nthInOrder(node.right, index, counter)

	//Implement your solution here
}
