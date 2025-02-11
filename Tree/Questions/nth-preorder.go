package questions

import "fmt"

func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
	//Implement your solution here
	if node == nil {
		return
	}
	(*counter) += 1
	if index == (*counter) {
		fmt.Print(node.value)
	}

	nthPreOrder(node.left, index, counter)
	nthPreOrder(node.right, index, counter)

}
