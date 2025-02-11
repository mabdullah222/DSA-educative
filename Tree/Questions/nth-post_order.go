package questions

import "fmt"

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	//Implement your solution here
	if node == nil {
		return
	}
	nthPostOrder(node.left, index, counter)
	nthPostOrder(node.right, index, counter)
	(*counter) += 1
	if index == (*counter) {
		fmt.Print(node.value)
	}
	//Implement your solution here
}
